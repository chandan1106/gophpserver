# Installation Guide for GoPHPServer

This guide provides detailed instructions for installing and configuring GoPHPServer on different platforms.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Installation Methods](#installation-methods)
   - [Pre-built Binaries](#pre-built-binaries)
   - [Build from Source](#build-from-source)
   - [Docker](#docker)
3. [Platform-specific Instructions](#platform-specific-instructions)
   - [Linux](#linux)
   - [macOS](#macos)
   - [Windows](#windows)
4. [Verifying Installation](#verifying-installation)
5. [Troubleshooting](#troubleshooting)

## Prerequisites

- **For running pre-built binaries**: None (standalone binary)
- **For PHP support**: PHP installed on your system
- **For building from source**: Go 1.16 or higher
- **For Docker**: Docker and Docker Compose

## Installation Methods

### Pre-built Binaries

1. Download the latest release for your platform from the [releases page](https://github.com/chandan1106/gophpserver/releases).
2. Extract the archive.
3. Make the binary executable (Linux/macOS only).
4. Run the server.

### Build from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/chandan1106/gophpserver.git
   cd gophpserver
   ```

2. Build the server:
   ```bash
   make build
   ```

3. The binary will be available in the `build` directory.

### Docker

1. Pull the Docker image:
   ```bash
   docker pull chandan1106/gophpserver:latest
   ```

2. Or build the image yourself:
   ```bash
   docker build -t gophpserver .
   ```

3. Run the container:
   ```bash
   docker run -d -p 8080:8080 -v $(pwd)/public:/app/public gophpserver
   ```

## Platform-specific Instructions

### Linux

1. Download the Linux binary:
   ```bash
   wget https://github.com/chandan1106/gophpserver/releases/download/v1.0.0/gophpserver-linux-amd64.tar.gz
   ```

2. Extract the archive:
   ```bash
   tar -xzf gophpserver-linux-amd64.tar.gz
   ```

3. Make the binary executable:
   ```bash
   chmod +x gophpserver
   ```

4. Run the server:
   ```bash
   ./gophpserver
   ```

5. (Optional) Install system-wide:
   ```bash
   sudo mv gophpserver /usr/local/bin/
   ```

### macOS

1. Download the macOS binary:
   ```bash
   curl -L -o gophpserver-darwin-amd64.tar.gz https://github.com/chandan1106/gophpserver/releases/download/v1.0.0/gophpserver-darwin-amd64.tar.gz
   ```

2. Extract the archive:
   ```bash
   tar -xzf gophpserver-darwin-amd64.tar.gz
   ```

3. Make the binary executable:
   ```bash
   chmod +x gophpserver
   ```

4. Run the server:
   ```bash
   ./gophpserver
   ```

5. (Optional) Install with Homebrew:
   ```bash
   brew tap chandan1106/gophpserver
   brew install gophpserver
   ```

### Windows

1. Download the Windows binary from the [releases page](https://github.com/chandan1106/gophpserver/releases).
2. Extract the zip file.
3. Run the executable:
   ```
   gophpserver.exe
   ```

4. (Optional) Add to PATH:
   - Right-click on "This PC" and select "Properties"
   - Click on "Advanced system settings"
   - Click on "Environment Variables"
   - Edit the "Path" variable and add the directory containing gophpserver.exe

## Verifying Installation

To verify that GoPHPServer is installed correctly:

1. Run the server:
   ```bash
   gophpserver
   ```

2. You should see output similar to:
   ```
   Starting server on 0.0.0.0:8080 with document root ./public
   ```

3. Open a web browser and navigate to:
   ```
   http://localhost:8080/
   ```

4. You should see the default page or your website content.

## Troubleshooting

### Common Issues

1. **"Permission denied" error**:
   - Make sure the binary is executable:
     ```bash
     chmod +x gophpserver
     ```

2. **"Address already in use" error**:
   - Another process is using port 8080. Use a different port:
     ```bash
     ./gophpserver -port 8081
     ```

3. **"PHP not found" error**:
   - Make sure PHP is installed and in your PATH, or specify the path:
     ```bash
     ./gophpserver -php /path/to/php
     ```

4. **Blank page when accessing PHP files**:
   - Check if PHP is installed correctly:
     ```bash
     php --version
     ```

For more help, please [open an issue](https://github.com/chandan1106/gophpserver/issues) on GitHub.