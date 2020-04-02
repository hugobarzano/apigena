#!/usr/bin/env bash

curl -X GET "http://localhost:3001/api/people" -H "accept: application/json"
curl -X GET "http://localhost:3002/api/people" -H "accept: application/json"
curl -X GET "http://localhost:3003/api/people" -H "accept: application/json"
