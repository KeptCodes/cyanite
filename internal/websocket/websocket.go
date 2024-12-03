package websocket

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/keptcodes/spongebob-desktop/internal/actions"
	"github.com/keptcodes/spongebob-desktop/internal/utils"
)

var upgrader = websocket.Upgrader{}

// StartServer starts both the HTTP server to serve files and the WebSocket server
func StartServer() {
    createFilesDirectory()
	// Start the HTTP server to serve files
	go startHTTPServer()

	// Start the WebSocket server
	startWebSocketServer()
}

// StartWebSocketServer starts the WebSocket server
func startWebSocketServer() {
	http.HandleFunc("/", handleConnection)
	server := &http.Server{
		Addr: ":8765",
	}

    var ip = utils.GetOutboundIP();

	log.Printf("WebSocket server started at ws://%s:8765", ip)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// handleConnection handles WebSocket connections
func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error establishing connection:", err)
		return
	}
	defer conn.Close()

	handleMessage(conn)
}


func handleMessage(conn *websocket.Conn) {
	// Loop to keep reading messages
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		fmt.Printf("Received: %s\n", p)

		// Call ProcessAction and get the response
		status, message := actions.ProcessAction(p)

		// Format the response
		response := fmt.Sprintf("{\"status\": \"%s\", \"message\": \"%s\"}", status, message)

		// Send the response back to the WebSocket client
		if err := conn.WriteMessage(messageType, []byte(response)); err != nil {
			log.Println("Error sending response:", err)
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
    var ip = utils.GetOutboundIP();

	log.Printf("HTTP server started at http://%s:8766", ip)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
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
		log.Printf("Created directory: %s", filesDir)
	} else {
		log.Printf("Directory %s already exists", filesDir)
	}
}
