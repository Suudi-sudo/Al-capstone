package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// handler function for the root endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to HTML
	w.Header().Set("Content-Type", "text/html")
	
	// Create a simple HTML response
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Go Web Server - Hello World</title>
		<style>
			body { 
				font-family: Arial, sans-serif; 
				max-width: 800px; 
				margin: 50px auto; 
				padding: 20px;
				background-color: #f5f5f5;
			}
			.container {
				background: white;
				padding: 30px;
				border-radius: 10px;
				box-shadow: 0 2px 10px rgba(0,0,0,0.1);
			}
			h1 { color: #00ADD8; }
			.info { 
				background: #e8f4f8; 
				padding: 15px; 
				border-radius: 5px; 
				margin: 20px 0;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1> Hello World from Go!</h1>
			<p>Congratulations! Your Go web server is running successfully.</p>
			<div class="info">
				<strong>Server Info:</strong><br>
				Time: %s<br>
				Method: %s<br>
				URL Path: %s<br>
				User Agent: %s
			</div>
			<p>This is a simple HTTP server built with Go's standard library.</p>
		</div>
	</body>
	</html>
	`
	
	// Format the HTML with dynamic content
	response := fmt.Sprintf(html, 
		time.Now().Format("2006-01-02 15:04:05"),
		r.Method,
		r.URL.Path,
		r.UserAgent(),
	)
	
	// Write the response
	fmt.Fprint(w, response)
}

// handler for the /api endpoint
func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := `{
		"message": "Hello world from Go API!",
		"timestamp": "` + time.Now().Format(time.RFC3339) + `",
		"status": "success"
	}`
	fmt.Fprint(w, response)
}

func main() {
	// Create a new HTTP multiplexer (router)
	mux := http.NewServeMux()
	
	// Register handlers
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/api", apiHandler)
	
	// Define the port
	port := ":8080"
	
	// Print startup message
	fmt.Printf("Go web server starting...\n")
	fmt.Printf(" Server running on http://localhost%s\n", port)
	fmt.Printf(" Visit http://localhost%s in your browser\n", port)
	fmt.Printf(" API endpoint: http://localhost%s/api\n", port)
	fmt.Printf("  Press Ctrl+C to stop the server\n\n")
	
	// Start the server
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
