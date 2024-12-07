package actions

import (
	"encoding/json"
	"fmt"
	"image/png"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
	"github.com/keptcodes/syra-server/internal/config"
	"github.com/keptcodes/syra-server/internal/utils"
)

type Action struct {
	SecretCode string        `json:"secret_code"`
	Action     string        `json:"action"`
	Inputs     []interface{} `json:"inputs"`
}

// Function to process the message and perform the action
func ProcessAction(message []byte) (string, string) {
	var action Action
	err := json.Unmarshal(message, &action)
	if err != nil {
		log.Println("Error parsing message:", err)
		return "error", "Error parsing message"
	}

	if !isValidSecretCode(action.SecretCode) {
		log.Println("Invalid secret code")
		return "error", "Invalid secret code"
	}

	// Based on the action type, you can perform different operations
	switch action.Action {
	case "key_press":
		// Handle key press action
		pressKeys(action.Inputs)
		return "success", "Key press action completed"
	case "screenshot":
		// Handle screenshot action
		status, message := takeScreenshot()
		return status, message
	case "shutdown":
		// Handle PC shutdown action
		status, message := shutdownPC()
		return status, message
	case "mouse_move":
		// Handle mouse movement action
		status, message := moveMouse(action.Inputs)
		return status, message
	case "mouse_click":
		// Handle mouse click action
		status, message := clickMouse(action.Inputs)
		return status, message
	// Add more actions as needed
	default:
		log.Printf("Unknown action: %s", action.Action)
		return "error", fmt.Sprintf("Unknown action: %s", action.Action)
	}
}

func isValidSecretCode(secretCode string) bool {
	// Read the saved secret code from config.json
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Println("Error reading config file:", err)
		return false
	}

	// Compare the incoming secret code with the one in config.json
	return cfg.SecretCode == secretCode
}

// TakeScreenshot captures a screenshot and saves it as a PNG file in the _data/files directory
func takeScreenshot() (string, string) {
	// Capture screenshot
	img, err := screenshot.CaptureDisplay(0)
	if err != nil {
		log.Println("Error taking screenshot:", err)
		return "error", "Error taking screenshot"
	}

	// Save the screenshot as a PNG file
	filePath := "_data/files/screenshot_" + time.Now().Format("20060102_150405") + ".png"
	urlPath := "files/screenshot_" + time.Now().Format("20060102_150405") + ".png"
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Error creating screenshot file:", err)
		return "error", "creating screenshot file"
	}
	defer file.Close()

	// Encode and save the image
	err = png.Encode(file, img)
	if err != nil {
		log.Println("Error encoding screenshot:", err)
		return "error", "Error encoding screenshot"
	}

	log.Printf("Screenshot saved as: %s", filePath)
	ip := utils.GetOutboundIP()
	imageURL := fmt.Sprintf("http://%s:8766/%s", ip, urlPath)
	return "success", imageURL
}

// PressShortcut simulates pressing a key combination
func pressKeys(keys []interface{}) error {
	// Convert keys to lowercase for consistency
	for i := range keys {
		key, ok := keys[i].(string)
		if !ok {
			return fmt.Errorf("invalid key type at index %d, expected string", i)
		}
		keys[i] = strings.ToLower(key)
	}

	// Press modifier keys first (e.g., shift, ctrl, alt)
	for _, key := range keys {
		keyStr, ok := key.(string)
		if !ok {
			return fmt.Errorf("invalid key type, expected string")
		}
		if isModifierKey(keyStr) {
			robotgo.KeyToggle(keyStr, "down") // Press the modifier key
		}
	}

	// Press other keys
	for _, key := range keys {
		keyStr, ok := key.(string)
		if !ok {
			return fmt.Errorf("invalid key type, expected string")
		}
		if !isModifierKey(keyStr) {
			robotgo.KeyTap(keyStr) // Tap the non-modifier key
		}
	}

	// Release modifier keys
	for _, key := range keys {
		keyStr, ok := key.(string)
		if !ok {
			return fmt.Errorf("invalid key type, expected string")
		}
		if isModifierKey(keyStr) {
			robotgo.KeyToggle(keyStr, "up") // Release the modifier key
		}
	}

	return nil
}


func isModifierKey(key string) bool {
	modifiers := []string{"shift", "ctrl", "alt", "command", "cmd"}
	for _, mod := range modifiers {
		if key == mod {
			return true
		}
	}
	return false
}

// ShutdownPC shuts down the PC
func shutdownPC() (string, string) {
	// Start the graceful shutdown process (without forcing)
	cmd := exec.Command("shutdown", "/s", "/t", "10") // Wait for 60 seconds before forcing shutdown
	err := cmd.Start()
	if err != nil {
		log.Println("Error initiating shutdown:", err)
		return "error", "Error initiating graceful shutdown"
	}

	log.Println("Graceful shutdown initiated. Waiting for 10 seconds...")

	// Wait for 10 seconds before checking if the shutdown is still in progress
	time.Sleep(10 * time.Second)

	// Now, force the shutdown if the system hasn't powered off
	log.Println("Force shutdown after waiting for 10 seconds...")
	cmd = exec.Command("shutdown", "/s", "/f", "/t", "0") // Forcibly shutdown the PC immediately
	err = cmd.Run()
	if err != nil {
		log.Println("Error forcing shutdown:", err)
		return "error", "Error forcing shutdown"
	}

	log.Println("PC is shutting down (forcefully if necessary)...")
	return "success", "PC is shutting down (forcefully if necessary)"
}

func moveMouse(inputs []interface{}) (string, string) {
	if len(inputs) < 2 {
		return "error", "Invalid inputs for mouse move"
	}

	// Ensure that inputs[0] and inputs[1] are numbers
	x, ok := inputs[0].(float64)
	if !ok {
		return "error", "Invalid X coordinate"
	}

	y, ok := inputs[1].(float64)
	if !ok {
		return "error", "Invalid Y coordinate"
	}

	robotgo.Move(int(x), int(y))
	return "success", fmt.Sprintf("Mouse moved to: (%d, %d)", int(x), int(y))
}

func clickMouse(inputs []interface{}) (string, string) {
	if len(inputs) < 1 {
		return "error", "Invalid input for mouse click"
	}

	// Ensure inputs[0] is a string
	clickType, ok := inputs[0].(string)
	if !ok {
		return "error", "Invalid click type"
	}

	switch clickType {
	case "left":
		robotgo.Click("left")
	case "right":
		robotgo.Click("right")
	case "middle":
		robotgo.Click("middle")
	default:
		return "error", "Unknown click type: " + clickType
	}

	return "success", fmt.Sprintf("%s click performed", clickType)
}
