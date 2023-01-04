#!/bin/bash

# Build the API app for linux
export ENV GOOS=linux
export ENV GOARCH=amd64
cd api && go build -o build/grocery-data-api && cd ..

# Build the web app
cd app && npm run build && cd ..
