package websocket

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/keptcodes/syra-server/internal/actions"
	"github.com/keptcodes/syra-server/internal/utils"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for testing purposes
	},
}

// StartServer starts both the HTTP server to serve files and the WebSocket server
func StartServer() {
	createFilesDirectory()
	// Start the HTTP server to serve files
	go startHTTPServer()

	// Start the WebSocket server with retry logic
	startWebSocketServerWithRetry()
}

// startWebSocketServerWithRetry starts the WebSocket server and retries if it fails
func startWebSocketServerWithRetry() {
	for {
		// Start the WebSocket server
		err := startWebSocketServer()
		if err != nil {
			log.Printf("WebSocket server failed to start: %v. Retrying in 5 seconds...\n", err)
			time.Sleep(5 * time.Second) // Wait before retrying
		}
	}
}

// startWebSocketServer starts the WebSocket server
func startWebSocketServer() error {
	http.HandleFunc("/", handleConnection)
	server := &http.Server{
		Addr: ":8765",
	}

	ip := utils.GetOutboundIP()
	log.Printf("WebSocket server started at ws://%s:8765", ip)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Error in WebSocket server: %v", err)
		return err
	}
	return nil
}

// handleConnection handles WebSocket connections
func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error establishing connection: %v", err)
		return
	}
	defer conn.Close()

	handleMessage(conn)
}

// handleMessage processes the message received from the WebSocket client
func handleMessage(conn *websocket.Conn) {
	// Loop to keep reading messages
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			return
		}
		fmt.Printf("Received: %s\n", p)

		// Call ProcessAction and get the response
		status, message := actions.ProcessAction(p)

		// Format the response
		response := fmt.Sprintf("{\"status\": \"%s\", \"message\": \"%s\"}", status, message)

		// Send the response back to the WebSocket client
		if err := conn.WriteMessage(messageType, []byte(response)); err != nil {
			log.Printf("Error sending response: %v", err)
			return
		}
	}
}

// Start HTTP server to serve files under _data/files directory
func startHTTPServer() {
	// Serve files from the _data/files directory
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("_data/files"))))

	// Start the HTTP server on port 8080
	server := &http.Server{
		Addr: ":8766",
	}
	ip := utils.GetOutboundIP()
	log.Printf("HTTP server started at http://%s:8766", ip)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Error in HTTP server: %v", err)
	}
}

func createFilesDirectory() {
	filesDir := "_data/files"

	// Check if the directory exists
	if _, err := os.Stat(filesDir); os.IsNotExist(err) {
		// If it doesn't exist, create the directory
		err := os.MkdirAll(filesDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating _data/files directory: %v", err)
		}
	}
}
