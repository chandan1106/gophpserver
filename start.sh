#!/bin/bash

# Kill any existing instances
pkill -f php_server || echo "No running instances found"

# Start the server
./php_server

# Note: This script will keep running until you press Ctrl+C