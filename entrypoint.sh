#!/bin/bash

# Start the nginx server
service nginx start

# Run the API application
cd /api && ./grocery-data-api
