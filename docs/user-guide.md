# GoPHPServer User Guide

This guide provides detailed instructions for installing, configuring, and using GoPHPServer.

## Table of Contents

1. [Overview](#overview)
2. [Installation](#installation)
3. [Basic Usage](#basic-usage)
4. [Configuration Options](#configuration-options)
5. [Directory Structure](#directory-structure)
6. [PHP Support](#php-support)
7. [Docker Deployment](#docker-deployment)
8. [Troubleshooting](#troubleshooting)

## Overview

GoPHPServer is a lightweight web server written in Go that can serve both static files and PHP scripts. It's designed to be a simpler alternative to the Nginx + PHP-FPM stack, with minimal configuration required.

## Installation

### Prerequisites

- Go 1.16 or higher (for building from source)
- PHP (for executing PHP scripts)

### Method 1: Using Pre-built Binaries

1. Download the pre-built binary for your platform from the [releases page](https://github.com/yourusername/gophpserver/releases).

2. Extract the archive:
   ```bash
   # For Linux/macOS
   tar -xzf gophpserver-1.0.0-linux-amd64.tar.gz
   
   # For Windows
   unzip gophpserver-1.0.0-windows-amd64.zip
   ```

3. Make the binary executable (Linux/macOS only):
   ```bash
   chmod +x gophpserver
   ```

### Method 2: Building from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/gophpserver.git
   cd gophpserver
   ```

2. Build the server:
   ```bash
   make build
   ```

3. The binary will be available in the `build` directory.

## Basic Usage

### Starting the Server

```bash
# Using the run script
./run.sh

# Or directly
./build/gophpserver
```

### Accessing Your Website

Open your web browser and navigate to:
```
http://localhost:8080/
```

### Stopping the Server

Press `Ctrl+C` in the terminal where the server is running, or use:
```bash
pkill -f gophpserver
```

## Configuration Options

GoPHPServer can be configured using command-line flags:

| Flag | Description | Default |
|------|-------------|---------|
| `-host` | Host to listen on | "0.0.0.0" |
| `-port` | Port to listen on | 8080 |
| `-docroot` | Document root directory | "./public" |
| `-php` | Path to PHP binary | auto-detect |
| `-log` | Log file path | stdout |

### Examples

```bash
# Change the port
./build/gophpserver -port 8000

# Specify a different document root
./build/gophpserver -docroot /var/www/html

# Specify a PHP binary path
./build/gophpserver -php /usr/bin/php

# Log to a file
./build/gophpserver -log ./logs/server.log
```

## Directory Structure

The recommended directory structure for your web application is:

```
/your-app/
├── gophpserver      # Server binary
└── public/          # Document root
    ├── index.php    # Default PHP page
    ├── index.html   # Default HTML page (fallback)
    └── ...          # Other web files
```

## PHP Support

GoPHPServer executes PHP scripts using the PHP interpreter installed on your system. It passes the appropriate environment variables to PHP to simulate a standard web server environment.

### Requirements

- PHP must be installed and available in your PATH
- PHP scripts must have the `.php` extension

### PHP Environment Variables

The following environment variables are set for PHP scripts:

- `REQUEST_METHOD`: The HTTP method used (GET, POST, etc.)
- `QUERY_STRING`: The query string from the URL
- `REQUEST_URI`: The full request URI
- `DOCUMENT_ROOT`: The document root directory
- `SCRIPT_FILENAME`: The full path to the PHP script
- `SERVER_SOFTWARE`: "GoPHPServer/1.0"
- `REMOTE_ADDR`: The client's IP address
- HTTP headers (prefixed with `HTTP_`)

## Docker Deployment

### Using Docker Compose

1. Make sure Docker and Docker Compose are installed on your system.

2. Start the server:
   ```bash
   docker-compose up -d
   ```

3. Stop the server:
   ```bash
   docker-compose down
   ```

### Using Docker Directly

1. Build the Docker image:
   ```bash
   docker build -t gophpserver .
   ```

2. Run the container:
   ```bash
   docker run -d -p 8080:8080 -v $(pwd)/public:/app/public --name gophpserver gophpserver
   ```

3. Stop the container:
   ```bash
   docker stop gophpserver
   ```

## Troubleshooting

### Common Issues

#### "PHP not found" error

**Problem**: The server cannot find the PHP interpreter.

**Solution**:
1. Make sure PHP is installed:
   ```bash
   php --version
   ```
2. If PHP is installed but not in the PATH, specify the path explicitly:
   ```bash
   ./build/gophpserver -php /path/to/php
   ```

#### "Address already in use" error

**Problem**: The port is already in use by another process.

**Solution**:
1. Use a different port:
   ```bash
   ./build/gophpserver -port 8081
   ```
2. Or stop the process using the current port:
   ```bash
   pkill -f gophpserver
   ```

#### PHP scripts not executing

**Problem**: PHP scripts are being downloaded instead of executed.

**Solution**:
1. Make sure the file has a `.php` extension
2. Check that PHP is installed and working:
   ```bash
   php -r "echo 'PHP is working';"
   ```

### Getting Help

If you encounter issues not covered in this guide, please:

1. Check the [GitHub issues](https://github.com/yourusername/gophpserver/issues) for similar problems
2. Open a new issue with detailed information about your problem