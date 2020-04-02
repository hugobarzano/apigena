#!/usr/bin/env bash


curl -X POST "http://localhost:3001/api/people" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"id\": \"33\", \"lname\": \"string\"}"
curl -X POST "http://localhost:3002/api/people" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"id\": \"33\", \"lname\": \"string\"}"
curl -X POST "http://localhost:3003/api/people" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"id\": \"33\", \"lname\": \"string\"}"