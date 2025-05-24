# GoPHPServer Quick Start Guide

This guide will help you get started with GoPHPServer in just a few minutes.

## Installation

### Option 1: Download Pre-built Binary

1. Download the binary for your platform from the [releases page](https://github.com/yourusername/gophpserver/releases).
2. Extract the archive and make the binary executable (if on Linux/macOS).

### Option 2: Build from Source

```bash
# Clone the repository
git clone https://github.com/yourusername/gophpserver.git
cd gophpserver

# Build the server
make build
```

## Running the Server

```bash
# Using the run script
./run.sh

# Or directly
./build/gophpserver
```

## Creating Your First PHP Website

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

## Basic Configuration

Change the port:
```bash
./run.sh -port 8000
```

Use a different document root:
```bash
./run.sh -docroot /path/to/your/website
```

## Docker Deployment

```bash
# Start with Docker Compose
docker-compose up -d

# Or build and run manually
docker build -t gophpserver .
docker run -d -p 8080:8080 -v $(pwd)/public:/app/public gophpserver
```

## Next Steps

- Read the full [User Guide](user-guide.md) for detailed information
- Check out the [README](../README.md) for an overview of features
- Explore the [examples](../examples) directory for sample configurations