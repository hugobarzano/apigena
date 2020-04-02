#!/usr/bin/env bash

curl -X PUT "http://localhost:3001/api/people/33" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"lname\": \"ss\"}"

curl -X PUT "http://localhost:3002/api/people/33" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"lname\": \"ss\"}"

curl -X PUT "http://localhost:3003/api/people/33" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"title\": \"ss\"}"

echo "FAIL"

#!/usr/bin/env bash

curl -X PUT "http://localhost:3001/api/people/ff" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"lname\": \"ss\"}"

curl -X PUT "http://localhost:3002/api/people/ff" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"lname\": \"ss\"}"

curl -X PUT "http://localhost:3003/api/people/ff" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"title\": \"ss\"}"