package gui

import (
	"fmt"
	"log"
	"os"

	"github.com/getlantern/systray"
	"github.com/keptcodes/cyanite-server/internal/config"
	"github.com/keptcodes/cyanite-server/internal/utils"
)

var secretCode string // Store the secret code

func Start() error {
	// Initialize config (create config if doesn't exist)
	if err := config.InitializeConfig(); err != nil {
		log.Fatal("Error initializing config: ", err)
	}

	// Read the secret code from the config
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal("Error reading config: ", err)
	}
	secretCode = cfg.SecretCode

	systray.Run(onReady, onExit)
	return nil
}

func onReady() {
	// Set up tray icon
	var iconData, err = os.ReadFile("assets/cyanite.ico")
	if err != nil {
		log.Fatal("Error loading icon: ", err)
	}
	systray.SetIcon(iconData)
	systray.SetTooltip("Cyanite Desktop Client")

	// Show current secret code in the tray
	systray.SetTooltip(fmt.Sprintf("Secret Code: %s", secretCode))

	// Add action to reset secret code
	mResetSecretCode := systray.AddMenuItem("Reset Secret Code", "Generate a new secret code")
	go func() {
		for range mResetSecretCode.ClickedCh {
			if err := config.ResetSecretCode(); err != nil {
				log.Println("Error resetting secret code:", err)
			} else {
				// After resetting, update the tooltip with the new secret code
				cfg, _ := config.ReadConfig()
				secretCode = cfg.SecretCode
				systray.SetTooltip(fmt.Sprintf("Secret Code: %s", secretCode))
			}
		}
	}()

	// Show Desktop and WebSocket IP
	desktopIP := utils.GetOutboundIP() // Function to get your desktop IP address  // Replace with actual WebSocket IP
	mIP := systray.AddMenuItem(fmt.Sprintf("Desktop IP: %s", desktopIP), fmt.Sprintf("Server IP: %s", desktopIP))
	go func() {
		for range mIP.ClickedCh {
			log.Printf("IP: %s", desktopIP)
		}
	}()

	// Quit Menu
	mQuit := systray.AddMenuItem("Quit", "Quit the application")
	go func() {
		for range mQuit.ClickedCh {
			systray.Quit()
		}
	}()
}

// Handle exit cleanup
func onExit() {
	// Clean up resources when the tray is closed
	log.Println("Tray exited")
}


