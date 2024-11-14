// Ethereum Transaction Parser
//
// Created by: Mohamed Riyad
// Date: 2024-11-14
// License: MIT
//

package server

import (
	"encoding/json"
	"ethereum-tx-parser/utils"
	"net/http"
	"testing"
)

const BASE_URL = "http://localhost:8080"

// TestServerHealth checks if the server is up and running before starting the actual tests
func TestServerHealth(t *testing.T) {
	// Perform a simple GET request to check if the server is up and running
	resp, err := http.Get(BASE_URL)
	if err != nil {
		// t.Fatalf("Failed to reach server: %v", err)
		t.Fatal("Failed to reach server. Is the server running?")
	}
	defer resp.Body.Close()

	// Check that the server responds with a 200 OK or 404 Not Found status
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code 200 or 404, got %d", resp.StatusCode)
	}
}

// TestHandleGetCurrentBlock tests the /block API endpoint
func TestHandleGetCurrentBlock(t *testing.T) {
	// Ensure the server is up and running
	TestServerHealth(t)

	// Send a GET request to the /block endpoint
	resp, err := http.Get(BASE_URL + "/block")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	// Check the response body
	var response map[string]int
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Ensure the current block is greater than 0 (assuming a valid blockchain)
	if response["currentBlock"] == 0 {
		t.Errorf("Expected current block to be greater than 0, got %d", response["currentBlock"])
	}
}

// TestHandleSubscribe tests the /subscribe API endpoint
func TestHandleSubscribe(t *testing.T) {
	// Ensure the server is up and running
	TestServerHealth(t)

	// Generate a random Ethereum address for each test run
	randomAddress := utils.GenerateRandomAddress()

	// Construct the URL with the address as a query parameter
	url := BASE_URL + "/subscribe?address=" + randomAddress

	// Perform the GET request using http.Get
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 201, got %d", resp.StatusCode)
	}

	// Check the response body
	var response map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response["message"] != "Subscribed successfully" {
		t.Errorf("Expected message 'Subscribed successfully', got %s", response["message"])
	}
}

// TestHandleGetTransactions tests the /transactions API endpoint
func TestHandleGetTransactions(t *testing.T) {
	// Ensure the server is up and running
	TestServerHealth(t)

	// Send a GET request to the /transactions endpoint with a query parameter
	resp, err := http.Get(BASE_URL + "/transactions?address=0x1234")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	// Check the response body
	var response struct {
		Transactions []string `json:"transactions"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// No error, empty or non-empty transactions array is fine
	if response.Transactions == nil {
		t.Fatalf("Expected 'transactions' field to be an array, got nil")
	}
}
