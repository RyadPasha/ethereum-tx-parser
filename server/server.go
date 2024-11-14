// Ethereum Transaction Parser
//
// Created by: Mohamed Riyad
// Date: 2024-11-14
// License: MIT
//

package server

import (
	"encoding/json"
	"log"
	"net/http"

	"ethereum-tx-parser/notification"
	"ethereum-tx-parser/parser"
	"ethereum-tx-parser/types"
)

// Server holds the dependencies required to run the server
type Server struct {
	parser *parser.Parser
	port   string
}

// NewServer initializes the server with a parser and port
func NewServer(parser *parser.Parser, port string) *Server {
	return &Server{
		parser: parser,
		port:   port,
	}
}

// =================================================================================================
// The following are the methods that can be called on the Server ==================================
// =================================================================================================

// handleGetCurrentBlock handles HTTP requests for retrieving the current block number.
func (s *Server) handleGetCurrentBlock(w http.ResponseWriter, r *http.Request) {
	block := s.parser.GetCurrentBlock()
	json.NewEncoder(w).Encode(map[string]int{"currentBlock": block})
}

// handleSubscribe handles HTTP requests to subscribe an address for transaction tracking.
func (s *Server) handleSubscribe(w http.ResponseWriter, r *http.Request) {
	// Get the address from the query parameters
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Missing address in query", http.StatusBadRequest)
		return
	}

	// Handle subscription
	if s.parser.Subscribe(address) {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Subscribed successfully"})
		notification.SendNotification("Subscribed successfully", address)
	} else {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{"message": "Already subscribed"})
	}
}

// handleGetTransactions handles HTTP requests to retrieve transactions for a subscribed address.
func (s *Server) handleGetTransactions(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	transactions := s.parser.GetTransactions(address)

	// Ensure transactions is an empty array if nil
	if transactions == nil {
		transactions = []types.Transaction{}
	}

	// Wrap transactions in a JSON object with the "transactions" key
	response := map[string]interface{}{
		"transactions": transactions,
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// =================================================================================================
// =================================================================================================
// =================================================================================================

// Start launches the HTTP server and registers the endpoint handlers.
func (s *Server) Start() {
	http.HandleFunc("/block", s.handleGetCurrentBlock)
	http.HandleFunc("/subscribe", s.handleSubscribe)
	http.HandleFunc("/transactions", s.handleGetTransactions)

	// Start the server
	log.Printf("Server is started on port %s...\n", s.port)
	err := http.ListenAndServe(":"+s.port, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
