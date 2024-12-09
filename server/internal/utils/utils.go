package utils

import (
	"log"
	"net"
	"os"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func CreateFilesDirectory(filesDir string) {


	// Check if the directory exists
	if _, err := os.Stat(filesDir); os.IsNotExist(err) {
		// If it doesn't exist, create the directory
		err := os.MkdirAll(filesDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating _data/files directory: %v", err)
		}
	}
}
