# GoPHPServer

A lightweight, high-performance web server written in Go that can serve both static files and PHP scripts. This server is designed to be a simpler alternative to the Nginx + PHP-FPM stack.

[![GitHub release](https://img.shields.io/github/v/release/chandan1106/gophpserver)](https://github.com/chandan1106/gophpserver/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## 📖 Documentation

Visit our [documentation site](https://chandan1106.github.io/gophpserver/) for complete installation and usage instructions.

## ✨ Features

- **Standalone Operation**: No external dependencies required (except PHP)
- **High Performance**: Uses Go's efficient HTTP server implementation
- **Security Features**: Includes security headers by default
- **PHP Support**: Execute PHP scripts directly using the PHP interpreter
- **Static File Serving**: Efficiently serves static content
- **Simple Configuration**: Easy to configure via command-line flags

## 🚀 Quick Start

### Installation

#### Option 1: Download Pre-built Binary

Download the binary for your platform from the [releases page](https://github.com/chandan1106/gophpserver/releases).

```bash
# Linux
wget https://github.com/chandan1106/gophpserver/releases/download/v1.0.0/gophpserver-linux-amd64.tar.gz
tar -xzf gophpserver-linux-amd64.tar.gz
chmod +x gophpserver

# macOS
wget https://github.com/chandan1106/gophpserver/releases/download/v1.0.0/gophpserver-darwin-amd64.tar.gz
tar -xzf gophpserver-darwin-amd64.tar.gz
chmod +x gophpserver

# Windows
# Download the zip file and extract it
```

#### Option 2: Build from Source

```bash
# Clone the repository
git clone https://github.com/chandan1106/gophpserver.git
cd gophpserver

# Build the server
make build
```

### Running the Server

```bash
# Start with default settings
./gophpserver

# Specify a different port and document root
./gophpserver -port 8000 -docroot /path/to/your/website
```

## 🔧 Configuration Options

| Flag | Description | Default |
|------|-------------|---------|
| `-host` | Host to listen on | "0.0.0.0" |
| `-port` | Port to listen on | 8080 |
| `-docroot` | Document root directory | "./public" |
| `-php` | Path to PHP binary | auto-detect |
| `-log` | Log file path | stdout |

## 🐳 Docker Support

```bash
# Start with Docker Compose
docker-compose up -d

# Or build and run manually
docker build -t gophpserver .
docker run -d -p 8080:8080 -v $(pwd)/public:/app/public gophpserver
```

## 📚 Documentation

- [Quick Start Guide](https://chandan1106.github.io/gophpserver/docs/quick-start.html)
- [User Guide](https://chandan1106.github.io/gophpserver/docs/user-guide.html)
- [Technical Documentation](https://chandan1106.github.io/gophpserver/docs/technical.html)
- [FAQ](https://chandan1106.github.io/gophpserver/docs/faq.html)

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Support

If you encounter any issues or have questions, please [open an issue](https://github.com/chandan1106/gophpserver/issues) on GitHub.