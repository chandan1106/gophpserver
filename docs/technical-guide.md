# GoPHPServer Technical Documentation

This document provides a technical overview of GoPHPServer's architecture, components, and functionality.

## Table of Contents

1. [Architecture Overview](#architecture-overview)
2. [Core Components](#core-components)
3. [Request Handling Flow](#request-handling-flow)
4. [PHP Execution Process](#php-execution-process)
5. [Security Features](#security-features)
6. [Performance Considerations](#performance-considerations)
7. [Extension Points](#extension-points)

## Architecture Overview

GoPHPServer is a lightweight web server written in Go that serves both static files and PHP scripts. It uses Go's standard library `net/http` package for HTTP handling and executes PHP scripts by invoking the PHP interpreter directly.

### High-Level Architecture

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│                 │     │                 │     │                 │
│  HTTP Request   │────▶│  Go HTTP Server │────▶│  Request Router │
│                 │     │                 │     │                 │
└─────────────────┘     └─────────────────┘     └────────┬────────┘
                                                         │
                                                         ▼
                        ┌─────────────────┐     ┌─────────────────┐
                        │                 │     │                 │
                        │   PHP Script    │◀────│  File Handler   │
                        │   Execution     │     │                 │
                        │                 │     └────────┬────────┘
                        └─────────────────┘              │
                                                         │
                                                         ▼
                                               ┌─────────────────┐
                                               │                 │
                                               │  Static File    │
                                               │    Serving      │
                                               │                 │
                                               └─────────────────┘
```

## Core Components

### 1. HTTP Server

The HTTP server is built using Go's `net/http` package. It handles incoming HTTP requests, manages connections, and routes requests to the appropriate handler.

```go
server := &http.Server{
    Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
    Handler:      createHandler(config),
    ReadTimeout:  time.Duration(config.Server.ReadTimeout) * time.Second,
    WriteTimeout: time.Duration(config.Server.WriteTimeout) * time.Second,
    IdleTimeout:  time.Duration(config.Server.IdleTimeout) * time.Second,
}
```

### 2. Request Router

The request router determines how to handle each incoming request based on the requested path and file type.

```go
func createHandler(config Config) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Add security headers
        // Determine file path
        // Check if file exists
        // Handle PHP files or serve static files
    })
}
```

### 3. PHP Handler

The PHP handler executes PHP scripts by invoking the PHP interpreter with the appropriate environment variables.

```go
func handlePHP(w http.ResponseWriter, r *http.Request, path string, config Config) {
    // Execute PHP directly
    cmd := exec.Command(config.PHPPath, path)
    
    // Set environment variables for PHP
    // Execute PHP and get output
    // Write the response
}
```

### 4. Static File Server

Static files are served using Go's built-in `http.ServeFile` function.

```go
// Serve static files
http.ServeFile(w, r, path)
```

## Request Handling Flow

1. **Request Reception**: The HTTP server receives an incoming request.
2. **Path Resolution**: The server determines the file path based on the request URL.
3. **File Type Detection**: The server checks if the requested file is a PHP script or a static file.
4. **Processing**:
   - For PHP files: The server executes the PHP interpreter with the script.
   - For static files: The server serves the file directly.
5. **Response**: The server sends the response back to the client.

### Detailed Flow Diagram

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│  Receive    │     │  Resolve    │     │  Check if   │     │ File Exists?│
│  HTTP       │────▶│  File       │────▶│  Path is    │────▶│  Yes/No     │
│  Request    │     │  Path       │     │  Directory  │     │             │
└─────────────┘     └─────────────┘     └─────────────┘     └──────┬──────┘
                                                                   │
                                                                   ▼
┌─────────────┐     ┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│  Send       │     │  Process    │     │  Determine  │     │  No: Return │
│  Response   │◀────│  File       │◀────│  File       │◀────│  404 Error  │
│  to Client  │     │  Content    │     │  Type       │     │             │
└─────────────┘     └─────────────┘     └─────────────┘     └─────────────┘
```

## PHP Execution Process

GoPHPServer executes PHP scripts by invoking the PHP interpreter directly with the script file. This approach is different from traditional PHP-FPM setups but provides a simpler architecture.

### Environment Variables

The following environment variables are set for PHP scripts:

- `REQUEST_METHOD`: The HTTP method used (GET, POST, etc.)
- `QUERY_STRING`: The query string from the URL
- `REQUEST_URI`: The full request URI
- `DOCUMENT_ROOT`: The document root directory
- `SCRIPT_FILENAME`: The full path to the PHP script
- `SERVER_SOFTWARE`: "GoPHPServer/1.0"
- `REMOTE_ADDR`: The client's IP address
- HTTP headers (prefixed with `HTTP_`)

### Execution Steps

1. Create a command to execute the PHP interpreter with the script file.
2. Set up environment variables to simulate a web server environment.
3. Execute the command and capture the output.
4. Send the output back to the client as the HTTP response.

```go
cmd := exec.Command(config.PHPPath, path)
cmd.Env = append(os.Environ(), environmentVariables...)
output, err := cmd.Output()
w.Write(output)
```

## Security Features

GoPHPServer includes several security features to help protect your applications:

### 1. Security Headers

The server adds the following security headers to all responses:

- `X-Content-Type-Options: nosniff`: Prevents MIME type sniffing.
- `X-Frame-Options: SAMEORIGIN`: Prevents clickjacking by restricting framing.
- `Content-Security-Policy: default-src 'self' 'unsafe-inline'`: Restricts resource loading.

### 2. Path Traversal Prevention

The server uses Go's `filepath.Join` function to safely join paths, which prevents directory traversal attacks.

```go
path := filepath.Join(config.DocRoot, r.URL.Path[1:])
```

### 3. Error Handling

The server includes proper error handling to prevent information leakage and ensure stability.

## Performance Considerations

### 1. Static File Serving

Static files are served directly using Go's efficient file serving capabilities, which include:
- Automatic content type detection
- Range request support
- Conditional GET support

### 2. PHP Execution

PHP scripts are executed on demand, which is efficient for low to medium traffic sites. For high-traffic sites, consider:
- Increasing the number of worker processes
- Using a reverse proxy like Nginx for caching
- Implementing application-level caching in your PHP code

### 3. Memory Usage

GoPHPServer has a small memory footprint since it doesn't maintain persistent PHP processes. Each PHP script execution creates a new process, which is efficient for low to medium traffic but may become a bottleneck for high-traffic sites.

## Extension Points

GoPHPServer can be extended in several ways:

### 1. Custom Middleware

You can add custom middleware to the HTTP server to implement features like:
- Authentication
- Rate limiting
- Request logging
- Response compression

### 2. Alternative PHP Execution Methods

The PHP execution mechanism can be modified to use different approaches:
- PHP-FPM via FastCGI
- PHP embedded in Go using CGO
- Alternative PHP implementations

### 3. Configuration Options

Additional configuration options can be added to customize the server's behavior:
- TLS/HTTPS support
- Virtual hosts
- URL rewriting
- Custom error pages

## Implementation Details

### Main Components

1. **Configuration**: Handles command-line flags and configuration options.
2. **HTTP Server**: Manages HTTP connections and request handling.
3. **File Handler**: Determines how to process different file types.
4. **PHP Executor**: Executes PHP scripts and captures output.

### Key Functions

- `main()`: Entry point that initializes the server.
- `createHandler()`: Creates the HTTP request handler.
- `handlePHP()`: Handles PHP script execution.

### Error Handling

GoPHPServer uses Go's error handling patterns to manage errors:
- File not found errors return 404 responses
- PHP execution errors return 500 responses
- Server startup errors are logged and cause the server to exit

## Conclusion

GoPHPServer provides a simple but effective way to serve PHP applications without the complexity of traditional web server setups. Its architecture is designed for simplicity and ease of use, making it ideal for development environments and small to medium production sites.