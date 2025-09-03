# Go Web Server Toolkit

A beginner-friendly toolkit for learning Go by building a simple web server.

##  What This Project Does

This project demonstrates how to build a basic HTTP web server in Go that:
- Serves HTML pages with dynamic content
- Provides a JSON API endpoint
- Uses Go's standard library (no external dependencies)
- Shows real-time server information

##  Prerequisites

Before you start, make sure you have:
- **Go installed** (version 1.19 or later)
- **A text editor** (VS Code, Sublime Text, or any editor)
- **A web browser** to test the server
- **Terminal/Command Prompt** access

##  Installation & Setup

### Step 1: Install Go
1. Visit [golang.org/dl](https://golang.org/dl/)
2. Download the installer for your operating system
3. Follow the installation instructions
4. Verify installation by running: `go version`

### Step 2: Get the Project
1. Download or clone this project
2. Navigate to the project directory in your terminal
3. You should see these files:
   - `main.go` (the main server code)
   - `go.mod` (Go module file)
   - `README.md` (this file)

### Step 3: Initialize the Go Module
\`\`\`bash
# If go.mod doesn't exist, run:
go mod init go-toolkit

# Download any dependencies (none needed for this project)
go mod tidy
\`\`\`

##  Running the Server

### Method 1: Direct Run (Recommended for beginners)
\`\`\`bash
go run main.go
\`\`\`

### Method 2: Build and Run
\`\`\`bash
# Build the executable
go build -o server main.go

# Run the executable
./server        # On macOS/Linux
server.exe      # On Windows
\`\`\`

##  Testing Your Server

Once the server is running, you'll see output like:
\`\`\`
 Go web server starting...
 Server running on http://localhost:8080
 Visit http://localhost:8080 in your browser
 API endpoint: http://localhost:8080/api
 Press Ctrl+C to stop the server
\`\`\`

### Test the Web Interface
1. Open your browser
2. Go to `http://localhost:8080`
3. You should see a styled "Hello World" page with server info

### Test the API Endpoint
1. Go to `http://localhost:8080/api`
2. You should see JSON response with timestamp

### Test with curl (optional)
\`\`\`bash
# Test main page
curl http://localhost:8080

# Test API endpoint
curl http://localhost:8080/api
\`\`\`

##  Stopping the Server

Press `Ctrl+C` in the terminal where the server is running.

##  Project Structure

\`\`\`
go-toolkit/
├── main.go          # Main server code
├── go.mod           # Go module definition
├── README.md        # Setup and usage instructions
├── TOOLKIT.md       # Comprehensive learning guide
├── AI_PROMPTS.md    # AI interaction log
└── .gitignore       # Git ignore file
\`\`\`

##  Customization Ideas

Try modifying the code to:
- Change the port number (line with `:8080`)
- Add new routes/endpoints
- Modify the HTML styling
- Add form handling
- Serve static files

##  Troubleshooting

**Port already in use?**
- Change `:8080` to `:8081` or another port in main.go

**Go command not found?**
- Make sure Go is properly installed and in your PATH

**Permission denied?**
- On some systems, you might need to use `sudo` or run as administrator

##  Next Steps

After getting this working:
1. Read through `TOOLKIT.md` for detailed explanations
2. Check `AI_PROMPTS.md` to see the learning process
3. Try the customization ideas above
4. Explore Go's documentation at [golang.org](https://golang.org)

##  License

This project is open source and available under the MIT License.
