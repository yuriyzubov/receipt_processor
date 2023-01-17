package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var regExpANum = regexp.MustCompile(`^[a-zA-Z0-9]*$`)

func isAlphanumeric(char byte) bool {
	return regExpANum.MatchString(string(char))
}

func alphanumPoints(retailer string) int {
	points := 0
	for i := 0; i < len(retailer); i++ {
		if isAlphanumeric(retailer[i]) {
			points++
		}
	}
	return points
}

func calculatePoints(receipt Receipt) int {
	points := 0
	log.Printf("%v", points)

	total, _ := strconv.ParseFloat(receipt.Total, 32)

	points += alphanumPoints(receipt.Retailer)

	log.Printf("Alphanumeric: %v", points)
	// every two items in the receipt
	points += pairPoints(receipt)

	log.Printf("Every two items: %v", points)

	// check if short description is a multiple of 3
	points += multipleOf3Points(receipt)

	log.Printf("Multiple of 3 points: %v", points)

	// check if total is a round number
	points += round_points(total)
	log.Printf("Check if round total: %v", points)

	// check if total is a multiple of 0.25
	points += round_quarter_points(total)
	log.Printf("Check if quarter-round total: %v", points)

	points += odd_day_points(receipt)
	log.Printf("Odd day points: %v", points)
	points += time_points(receipt)

	//check time points
	log.Printf("Timeframe points: %v", points)
	return points
}

func pairPoints(receipt Receipt) int {
	return int(math.Floor(float64(len(receipt.Items)/2))) * 5
}

func generateId() string {
	id := uuid.New()
	return id.String()
}

func multipleOf3Points(receipt Receipt) int {
	points := 0
	for i := 0; i < len(receipt.Items); i++ {
		itemPrice, err := strconv.ParseFloat(receipt.Items[i].Price, 32)
		if err != nil {
			continue
		}
		if len(strings.TrimSpace(receipt.Items[i].ShortDescription))%3 == 0 {
			itemPrice = math.Ceil(itemPrice * 0.2)
			points += int(itemPrice)
		}
	}
	return points
}

func odd_day_points(receipt Receipt) int {
	YYMMDD := "2006-01-02"
	date, err := time.Parse(YYMMDD, receipt.PurchaseDate)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	if (date.Day() % 2) != 0 {
		return 6
	}
	return 0
}

func time_points(receipt Receipt) int {
	dateString := string(receipt.PurchaseTime) + ":00"
	HHMMSS24h := "15:04:05"

	date, err := time.Parse(HHMMSS24h, dateString)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	beforeTime, err := time.Parse(HHMMSS24h, "14:00:00")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	afterTime, err := time.Parse(HHMMSS24h, "16:00:00")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	if date.After(beforeTime) && date.Before(afterTime) {
		return 10
	}
	return 0
}

func round_points(total float64) int {
	if int(total*100)%100 == 0 {
		return 50
	}
	return 0
}

func round_quarter_points(total float64) int {
	if int(total*100)%25 == 0 {
		return 25
	}
	return 0
}
