#!/bin/bash

echo "1. Testing Get All Customers"
curl -X GET http://localhost:8080/customers \
  -H "Content-Type: application/json" | json_pp

echo -e "\n\n2. Testing Get Customer by ID (ID: 1)"
curl -X GET http://localhost:8080/customers/1 \
  -H "Content-Type: application/json" | json_pp

echo -e "\n\n3. Testing Create New Customer"
curl -X POST http://localhost:8080/customers \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "phone": 1234567890
  }' | json_pp

echo -e "\n\n4. Verify New Customer Added (Get All)"
curl -X GET http://localhost:8080/customers \
  -H "Content-Type: application/json" | json_pp