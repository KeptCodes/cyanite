package config

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/keptcodes/cyanite-server/internal/utils"
)

// SecretCode structure for config.json
type Config struct {
	SecretCode string `json:"secret_code"`
}

// Default config file path
const configFilePath = "_data/config.json"

// Initialize the random seed
func init() {

	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// ReadConfig reads the config file and returns the Config structure
func ReadConfig() (*Config, error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("error decoding config file: %v", err)
	}
	return &config, nil
}

// WriteConfig writes the Config structure to the config.json file
func WriteConfig(config *Config) error {
	file, err := os.Create(configFilePath)
	if err != nil {
		return fmt.Errorf("error creating config file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("error encoding config file: %v", err)
	}
	return nil
}

// GenerateRandomSecretCode generates a 5-character random string
func GenerateRandomSecretCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	secretCode := make([]byte, 5)
	for i := range secretCode {
		secretCode[i] = charset[rand.Intn(len(charset))]
	}
	return string(secretCode)
}

// InitializeConfig initializes the config file with a random secret code if it doesn't exist
func InitializeConfig() error {
	// Check if the config file exists
	utils.CreateFilesDirectory("_data")
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		// Generate a random secret code
		randomSecretCode := GenerateRandomSecretCode()

		// Create a new config with the generated secret code
		config := &Config{SecretCode: randomSecretCode}

		// Write the config to the file
		if err := WriteConfig(config); err != nil {
			return fmt.Errorf("error creating config file: %v", err)
		}

		log.Printf("Config file created with secret code: %s", randomSecretCode)
	}
	return nil
}

// ResetSecretCode generates a new random secret code and updates the config
func ResetSecretCode() error {
	newSecretCode := GenerateRandomSecretCode()
	config := &Config{SecretCode: newSecretCode}

	if err := WriteConfig(config); err != nil {
		return fmt.Errorf("error updating secret code: %v", err)
	}

	log.Printf("Secret code reset to: %s", newSecretCode)
	return nil
}
