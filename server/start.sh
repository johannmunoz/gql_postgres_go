#!/bin/bash

function cleanup {
    kill "$ACCOUNTS_PID"
    kill "$PRODUCTS_PID"
    kill "$REVIEWS_PID"
}
trap cleanup EXIT

go build -o /tmp/srv-accounts ./cmd/accounts
go build -o /tmp/srv-products ./cmd/products
go build -o /tmp/srv-reviews ./cmd/reviews

/tmp/srv-accounts &
ACCOUNTS_PID=$!

/tmp/srv-products &
PRODUCTS_PID=$!

/tmp/srv-reviews &
REVIEWS_PID=$!

sleep 1

node gateway/index.js
