package main

import (
	"log"

	"github.com/keptcodes/syra-server/internal/gui"
	"github.com/keptcodes/syra-server/internal/websocket"
)

func main() {
    // Start the WebSocket server in a separate goroutine
    go websocket.StartServer()

    // Start the GUI (System Tray or Window)
    if err := gui.Start(); err != nil {
        log.Fatal(err)
    }
}
