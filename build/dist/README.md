# GoPHPServer

A lightweight, high-performance web server written in Go that can serve both static files and PHP scripts. This server is designed to be a simpler alternative to the Nginx + PHP-FPM stack.

## Features

- **Standalone Operation**: No external dependencies required (except PHP)
- **High Performance**: Uses Go's efficient HTTP server implementation
- **Security Features**: Includes security headers by default
- **PHP Support**: Execute PHP scripts directly using the PHP interpreter
- **Static File Serving**: Efficiently serves static content
- **Simple Configuration**: Easy to configure via command-line flags

## Requirements

- Go 1.16 or higher (for building from source)
- PHP (for executing PHP scripts)

## Installation

### From Source

1. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/gophpserver.git
   cd gophpserver
   ```

2. Build the server:
   ```bash
   make build
   ```

3. The binary will be available in the `build` directory.

### Pre-built Binaries

Download the pre-built binary for your platform from the [releases page](https://github.com/yourusername/gophpserver/releases).

## Usage

```bash
./gophpserver [options]
```

### Command-line Options

- `-host`: Host to listen on (default: "0.0.0.0")
- `-port`: Port to listen on (default: 8080)
- `-docroot`: Document root directory (default: "./public")
- `-php`: Path to PHP binary (default: auto-detect)
- `-log`: Log file path (default: stdout)

### Examples

```bash
# Start with default settings
./gophpserver

# Specify a different port and document root
./gophpserver -port 8000 -docroot /var/www/html

# Specify a PHP binary path
./gophpserver -php /usr/bin/php

# Log to a file
./gophpserver -log /var/log/gophpserver.log
```

## Directory Structure

For a basic setup, create the following directory structure:

```
/your-app/
├── gophpserver      # Server binary
└── public/          # Document root
    ├── index.php    # Default PHP page
    ├── index.html   # Default HTML page (fallback)
    └── ...          # Other web files
```

## Docker Support

### Using Docker

```bash
# Build the Docker image
docker build -t gophpserver .

# Run the container
docker run -d -p 8080:8080 -v $(pwd)/public:/app/public --name gophpserver gophpserver
```

### Using Docker Compose

```bash
# Start the server
docker-compose up -d

# Stop the server
docker-compose down
```

## Development

### Building for Multiple Platforms

```bash
make build-all
```

### Creating Distribution Packages

```bash
make dist
```

## License

MIT