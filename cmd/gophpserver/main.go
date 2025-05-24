package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Configuration options
type Config struct {
	Host      string
	Port      int
	DocRoot   string
	PHPPath   string
	LogFile   string
}

func main() {
	// Parse command line flags
	config := Config{}
	flag.StringVar(&config.Host, "host", "0.0.0.0", "Host to listen on")
	flag.IntVar(&config.Port, "port", 8080, "Port to listen on")
	flag.StringVar(&config.DocRoot, "docroot", "./public", "Document root directory")
	flag.StringVar(&config.PHPPath, "php", "", "Path to PHP binary (default: auto-detect)")
	flag.StringVar(&config.LogFile, "log", "", "Log file path (default: stdout)")
	flag.Parse()

	// Setup logging
	if config.LogFile != "" {
		logFile, err := os.OpenFile(config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		defer logFile.Close()
		log.SetOutput(logFile)
	}

	// Ensure document root exists
	if _, err := os.Stat(config.DocRoot); os.IsNotExist(err) {
		if err := os.MkdirAll(config.DocRoot, 0755); err != nil {
			log.Fatalf("Failed to create document root: %v", err)
		}
	}

	// Find PHP binary if not specified
	if config.PHPPath == "" {
		phpPath, err := exec.LookPath("php")
		if err != nil {
			log.Fatalf("PHP not found. Please install PHP or specify path with -php flag")
		}
		config.PHPPath = phpPath
	} else {
		// Verify the specified PHP path
		if _, err := os.Stat(config.PHPPath); os.IsNotExist(err) {
			log.Fatalf("PHP binary not found at %s", config.PHPPath)
		}
	}
	log.Printf("Using PHP at: %s", config.PHPPath)

	// Create HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler: createHandler(config),
	}

	// Start server
	log.Printf("Starting server on http://%s:%d with document root %s", config.Host, config.Port, config.DocRoot)
	log.Fatal(server.ListenAndServe())
}

func createHandler(config Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add security headers
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "SAMEORIGIN")
		w.Header().Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline'")

		// Get the file path
		path := filepath.Join(config.DocRoot, r.URL.Path[1:])
		if path == filepath.Join(config.DocRoot, "") {
			// Try index.php first, then index.html
			phpIndex := filepath.Join(config.DocRoot, "index.php")
			htmlIndex := filepath.Join(config.DocRoot, "index.html")
			
			if _, err := os.Stat(phpIndex); err == nil {
				path = phpIndex
			} else if _, err := os.Stat(htmlIndex); err == nil {
				path = htmlIndex
			} else {
				path = phpIndex // Default to index.php even if it doesn't exist
			}
		}

		// Check if the file exists
		info, err := os.Stat(path)
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error accessing %s: %v", path, err)
			return
		}

		// If it's a directory, look for index.php or index.html
		if info.IsDir() {
			phpIndex := filepath.Join(path, "index.php")
			htmlIndex := filepath.Join(path, "index.html")
			
			if _, err := os.Stat(phpIndex); err == nil {
				path = phpIndex
			} else if _, err := os.Stat(htmlIndex); err == nil {
				path = htmlIndex
			} else {
				http.NotFound(w, r)
				return
			}
		}

		// Handle PHP files
		if strings.HasSuffix(path, ".php") {
			handlePHP(w, r, path, config)
			return
		}

		// Serve static files
		http.ServeFile(w, r, path)
	})
}

func handlePHP(w http.ResponseWriter, r *http.Request, path string, config Config) {
	// Execute PHP directly
	cmd := exec.Command(config.PHPPath, path)
	
	// Set environment variables for PHP
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("REQUEST_METHOD=%s", r.Method),
		fmt.Sprintf("QUERY_STRING=%s", r.URL.RawQuery),
		fmt.Sprintf("REQUEST_URI=%s", r.URL.RequestURI()),
		fmt.Sprintf("DOCUMENT_ROOT=%s", config.DocRoot),
		fmt.Sprintf("SCRIPT_FILENAME=%s", path),
		fmt.Sprintf("SERVER_SOFTWARE=GoPHPServer/1.0"),
		fmt.Sprintf("REMOTE_ADDR=%s", strings.Split(r.RemoteAddr, ":")[0]),
	)

	// Add HTTP headers to environment
	for header, values := range r.Header {
		headerName := fmt.Sprintf("HTTP_%s", strings.ReplaceAll(strings.ToUpper(header), "-", "_"))
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", headerName, values[0]))
	}

	// Execute PHP and get output
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, "PHP Execution Error", http.StatusInternalServerError)
		log.Printf("Error executing PHP: %v", err)
		return
	}

	// Set content type to HTML
	w.Header().Set("Content-Type", "text/html")
	
	// Write the response
	w.Write(output)
}