package main

import (
	"log"

	"github.com/keptcodes/spongebob-desktop/internal/gui"
	"github.com/keptcodes/spongebob-desktop/internal/websocket"
)

func main() {
    // Start the WebSocket server in a separate goroutine
    go websocket.StartServer()

    // Start the GUI (System Tray or Window)
    if err := gui.Start(); err != nil {
        log.Fatal(err)
    }
}
