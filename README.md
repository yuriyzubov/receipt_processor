# receipt_processor
This web service calculates the number of points that should be rewarded for a receipt according to
the following rules:
* One point for every alphanumeric character in the retailer name.
* 50 points if the total is a round dollar amount with no cents.
* 25 points if the total is a multiple of `0.25`.
* 5 points for every two items on the receipt.
* If the trimmed length of the item description is a multiple of 3, multiply the price by `0.2` and round up to the nearest integer. 
The result is the number of points earned.
* 6 points if the day in the purchase date is odd.
* 10 points if the time of purchase is after 2:00pm and before 4:00pm.


To run the web service locally, please install *docker* and *golang*.
1. Install docker: https://docs.docker.com/get-docker/
   
   Install golang: https://go.dev/doc/install


2. In a terminal, switch to the receipt_processor directory:
    ```
   cd PATH_TO_DIRECTORY/receipt_processor
   ```

4. Run the docker-compose.yml file:
   ```
   docker compose up --build 
   ```

5. The simplest way to test the web service could be achieved by sending
   POST and GET HTTP methods through the Postman application. The examples of JSON Payloads could be found
   in the receipt_processor/examples folder.
   ```
   POST HTTP methods + http://localhost:8080/receipts/process + raw JSON data-> JSON id data:
   {"id" : "9fb03564-47a5-4ca6-b14d-6dddbbe26b1b"}
   ```
   
   ```
   GET request + http://localhost:8080/receipts/9fb03564-47a5-4ca6-b14d-6dddbbe26b1b/points
   + raw JSON data-> JSON *points* data:
     { "points" : 109 }
   ```
   
