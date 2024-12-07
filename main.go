package main

import (
	"log"
	"os"

	"github.com/keptcodes/syra-server/internal/gui"
	"github.com/keptcodes/syra-server/internal/websocket"
)

func init() {
	// Set up error logging to a file
	file, err := os.OpenFile("syra_logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	log.SetOutput(file)
	log.Println("Starting Syra application...")

	if err != nil {
		log.Fatalf("Failed to add to startup: %v", err)
	}
}



func main() {
	// Start the WebSocket server in a separate goroutine
	go websocket.StartServer()

	// Start the GUI (System Tray or Window)
	if err := gui.Start(); err != nil {
		log.Fatalf("Failed to start GUI: %v", err)
	}
}
