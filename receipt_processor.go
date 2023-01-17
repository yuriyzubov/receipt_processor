package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"sync"
)

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
}

type responseId struct {
	Id string `json:"id"`
}

type responsePoints struct {
	Points int `json:"points""`
}

type Controller struct {
	mu sync.RWMutex
	m  map[string]Receipt
}

func (c *Controller) AddReceipt(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	var receipt Receipt
	err = json.Unmarshal(body, &receipt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	id := generateId()
	idJson, err := json.Marshal(responseId{id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	c.mu.Lock()
	c.m[id] = receipt
	c.mu.Unlock()
	w.Write(idJson)

}

func (c *Controller) GetPoints(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Receipt not found."))
		return
	}
	receipt := c.m[id]
	points := calculatePoints(receipt)
	pointsJson, err := json.Marshal(responsePoints{points})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(pointsJson)

}

func (c *Controller) Start() {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", c.AddReceipt).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", c.GetPoints).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
func main() {

	c := &Controller{
		m: make(map[string]Receipt),
	}
	c.Start()
}
