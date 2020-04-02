#!/usr/bin/env bash

curl -X GET "http://localhost:3001/api/people/33" -H "accept: application/json"
curl -X GET "http://localhost:3002/api/people/33" -H "accept: application/json"
curl -X GET "http://localhost:3003/api/people/33" -H "accept: application/json"


echo "FAIL"

curl -X GET "http://localhost:3001/api/people/ff" -H "accept: application/json"
curl -X GET "http://localhost:3002/api/people/ff" -H "accept: application/json"
curl -X GET "http://localhost:3003/api/people/ff" -H "accept: application/json"