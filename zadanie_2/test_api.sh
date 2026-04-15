#!/bin/bash

BASE_URL="http://localhost:8000/api/products"

echo "TESTING PRODUCTS ENDPOINTS"
echo ""


echo "TEST 1 (POST)"
curl -s -X POST "$BASE_URL" \
     -H "Content-Type: application/json" \
     -d '{"name": "Kremówka", "price": 21.36, "description": "super jest ta kremówka"}'
echo -e "\n\n"


echo "TEST 2 (GET all)"
curl -s -X GET "$BASE_URL"
echo -e "\n\n"


ID=1

echo "TEST 3 (GET id=$ID)"
curl -s -X GET "$BASE_URL/$ID"
echo -e "\n\n"


echo "TEST 4 (PUT id=$ID)"
curl -s -X PUT "$BASE_URL/$ID" \
     -H "Content-Type: application/json" \
     -d '{"price": 21.37}'
echo -e "\n"
curl -s -X GET "$BASE_URL/$ID"
echo -e "\n\n"

echo "TEST 5 (DELETE id=$ID)"
curl -s -X DELETE "$BASE_URL/$ID"
echo -e "\n"
curl -s -X GET "$BASE_URL/$ID"
echo -e "\n\n"
