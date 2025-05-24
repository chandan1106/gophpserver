# GoPHPServer

A lightweight, high-performance web server written in Go that can serve both static files and PHP scripts. This server is designed to be a simpler alternative to the Nginx + PHP-FPM stack.

## Features

- **Standalone Operation**: No external dependencies required (except PHP)
- **High Performance**: Uses Go's efficient HTTP server implementation
- **Security Features**: Includes security headers by default
- **PHP Support**: Execute PHP scripts directly using the PHP interpreter
- **Static File Serving**: Efficiently serves static content
- **Simple Configuration**: Easy to configure via command-line flags

## Quick Start

### Installation

#### Option 1: Download Pre-built Binary

Download the binary for your platform from the [releases page](https://github.com/yourusername/gophpserver/releases).

#### Option 2: Build from Source

```bash
# Clone the repository
git clone https://github.com/yourusername/gophpserver.git
cd gophpserver

# Build the server
make build
```

### Running the Server

```bash
# Using the run script
./run.sh

# Or directly
./build/gophpserver
```

### Creating Your First PHP Website

1. Create a directory for your website files:
   ```bash
   mkdir -p public
   ```

2. Create a simple PHP file:
   ```bash
   echo '<?php echo "<h1>Hello from GoPHPServer!</h1>"; ?>' > public/index.php
   ```

3. Start the server:
   ```bash
   ./run.sh
   ```

4. Open your browser and navigate to:
   ```
   http://localhost:8080/
   ```

## Documentation

- [Quick Start Guide](docs/quick-start.md) - Get up and running in minutes
- [User Guide](docs/user-guide.md) - Detailed usage instructions
- [FAQ](docs/faq.md) - Frequently asked questions
- [Examples](examples) - Example applications

## Command-line Options

- `-host`: Host to listen on (default: "0.0.0.0")
- `-port`: Port to listen on (default: 8080)
- `-docroot`: Document root directory (default: "./public")
- `-php`: Path to PHP binary (default: auto-detect)
- `-log`: Log file path (default: stdout)

## Docker Support

### Using Docker Compose

```bash
# Start the server
docker-compose up -d

# Stop the server
docker-compose down
```

### Using Docker Directly

```bash
# Build the Docker image
docker build -t gophpserver .

# Run the container
docker run -d -p 8080:8080 -v $(pwd)/public:/app/public --name gophpserver gophpserver
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