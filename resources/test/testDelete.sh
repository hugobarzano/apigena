#!/usr/bin/env bash
curl -X DELETE "http://localhost:3001/api/people/33" -H "accept: application/json"
curl -X DELETE "http://localhost:3002/api/people/33" -H "accept: application/json"
curl -X DELETE "http://localhost:3003/api/people/33" -H "accept: application/json"