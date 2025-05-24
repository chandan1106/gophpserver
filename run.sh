#!/bin/bash

# Kill any existing instances
pkill -f gophpserver || echo "No running instances found"

# Start the server
./build/gophpserver "$@"