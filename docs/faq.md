# Frequently Asked Questions (FAQ)

## General Questions

### What is GoPHPServer?

GoPHPServer is a lightweight web server written in Go that can serve both static files and PHP scripts. It's designed to be a simpler alternative to the Nginx + PHP-FPM stack, with minimal configuration required.

### Why use GoPHPServer instead of Apache or Nginx?

GoPHPServer is much simpler to set up and use, especially for development environments or small projects. It's a single binary with no complex configuration files, making it ideal for quick PHP development or testing.

### What are the system requirements?

- Any operating system that Go supports (Windows, macOS, Linux)
- PHP installed on the system (for executing PHP scripts)
- Minimal memory and CPU requirements

## Installation Questions

### Do I need to install Go to use GoPHPServer?

No, you can download the pre-built binary for your platform from the releases page. You only need Go installed if you want to build from source.

### How do I update to a newer version?

Simply download the latest binary and replace your existing one. Your configuration and website files will not be affected.

### Can I install GoPHPServer globally on my system?

Yes, you can move the binary to a directory in your PATH, such as `/usr/local/bin` on Linux/macOS or add it to your PATH on Windows.

## Configuration Questions

### How do I change the port?

Use the `-port` flag:
```bash
./gophpserver -port 8000
```

### Can I serve multiple websites?

Currently, GoPHPServer serves a single document root. To serve multiple websites, you would need to run multiple instances on different ports.

### Does GoPHPServer support HTTPS?

Not directly in the current version. For HTTPS support, you can use a reverse proxy like Nginx or a service like Cloudflare.

## PHP Questions

### Which PHP versions are supported?

GoPHPServer should work with any modern PHP version (5.6 and above).

### How does GoPHPServer execute PHP scripts?

GoPHPServer executes PHP scripts by running the PHP interpreter directly with the script file. It sets up the appropriate environment variables to simulate a standard web server environment.

### Can I use PHP frameworks like Laravel or WordPress?

Yes, but some advanced frameworks might require additional configuration. Simple frameworks and CMS systems should work out of the box.

## Performance Questions

### How does GoPHPServer perform compared to Nginx + PHP-FPM?

GoPHPServer is designed for simplicity rather than maximum performance. For most small to medium websites, the performance difference will be negligible. For high-traffic production sites, Nginx + PHP-FPM might offer better performance and more advanced features.

### Is GoPHPServer suitable for production use?

GoPHPServer can be used in production for small to medium websites with moderate traffic. For high-traffic sites, consider using it behind a more robust web server like Nginx.

## Troubleshooting Questions

### PHP scripts are being downloaded instead of executed

Make sure PHP is installed and in your PATH, or specify the path to PHP using the `-php` flag.

### I'm getting "permission denied" errors

Make sure the GoPHPServer binary is executable and that it has permission to read your website files.

### How do I report a bug?

Please open an issue on the GitHub repository with detailed information about the bug, including your operating system, PHP version, and steps to reproduce the issue.