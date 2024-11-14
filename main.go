// Ethereum Transaction Parser
//
// Created by: Mohamed Riyad
// Date: 2024-11-14
// License: MIT
//

package main

import (
	"log"
	"time"

	"ethereum-tx-parser/parser"
	"ethereum-tx-parser/server"
	"ethereum-tx-parser/utils"
)

func main() {
	port := utils.GetPort()
	log.Println("Starting Ethereum Tx Parser...")

	// Initialize the parser
	parser := parser.NewParser()

	// Start the block updater to keep the current block number up to date
	parser.StartBlockUpdater(10 * time.Second) // Update every 10 seconds

	// Initialize the server with the custom port
	srv := server.NewServer(parser, port)

	// Start the server
	srv.Start()
}
