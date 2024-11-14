// Ethereum Transaction Parser
//
// Created by: Mohamed Riyad
// Date: 2024-11-14
// License: MIT
//

package utils

import (
	"crypto/rand"
	"encoding/hex"
	"os"

	"ethereum-tx-parser/config"
)

// Helper function to generate a random Ethereum address
func GenerateRandomAddress() string {
	// Generate 20 random bytes (Ethereum addresses are 20 bytes long)
	randomBytes := make([]byte, 20)
	_, err := rand.Read(randomBytes)
	if err != nil {
		// Fallback in case of error (shouldn't happen)
		return "0x0000000000000000000000000000000000000000"
	}

	// Convert the random bytes to a hexadecimal string and prefix with '0x'
	return "0x" + hex.EncodeToString(randomBytes)
}

// Retrieves the value of an environment variable or returns a default value.
func GetEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Returns the port from the environment variable "PORT", or defaults to config.DefaultAppPort.
func GetPort() string {
	return GetEnv("PORT", config.DefaultAppPort)
}
