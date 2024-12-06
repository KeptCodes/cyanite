package gui

import (
	"fmt"
	"log"
	"os"

	"github.com/getlantern/systray"
	"github.com/keptcodes/syra-server/internal/config"
	"github.com/keptcodes/syra-server/internal/utils"
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
	var iconData, err = os.ReadFile("assets/syra.ico")
	if err != nil {
		log.Fatal("Error loading icon: ", err)
	}
	systray.SetIcon(iconData)
	systray.SetTooltip("syra Desktop Client")

	// Show current secret code in the tray
	systray.SetTooltip(fmt.Sprintf("Secret Code: %s", secretCode))

	// Add menu to check mobile connection status
	mMobileStatus := systray.AddMenuItem("Check Mobile Connection", "Check if mobile is connected")
	go func() {
		for range mMobileStatus.ClickedCh {
			// Call a function to check if mobile is connected
			ip, connected := getMobileIP() // Replace with your actual function to get the mobile IP
			if connected {
				log.Printf("Mobile is connected with IP: %s", ip)
				systray.SetTooltip(fmt.Sprintf("Mobile IP: %s", ip))
			} else {
				log.Println("Mobile is not connected")
				systray.SetTooltip("Mobile is not connected")
			}
		}
	}()

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
	desktopIP := utils.GetOutboundIP() // Function to get your desktop IP address
	wsIP := getWebSocketIP()     // Replace with actual WebSocket IP
	mIP := systray.AddMenuItem(fmt.Sprintf("Desktop IP: %s", desktopIP), "WebSocket Server IP: "+wsIP)
	go func() {
		for range mIP.ClickedCh {
			log.Printf("Desktop IP: %s", desktopIP)
			log.Printf("WebSocket IP: %s", wsIP)
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

// Function to get WebSocket IP (example: return server's IP)
func getWebSocketIP() string {
	// Replace with the actual WebSocket server IP
	return "127.0.0.1" // Localhost for example
}

// Function to get mobile IP
func getMobileIP() (string, bool) {
	// Replace with the actual logic to check if mobile is connected
	// For now, we simulate a connected mobile device
	return "192.168.1.101", true
}
