// Ethereum Transaction Parser
//
// Created by: Mohamed Riyad
// Date: 2024-11-14
// License: MIT
//

package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"ethereum-tx-parser/config"
	"ethereum-tx-parser/memory"
	"ethereum-tx-parser/types"
)

// Parser implements the Parser interface, handling Ethereum blockchain data retrieval.
type Parser struct {
	currentBlock  int                   // Last parsed Ethereum block number
	memoryManager *memory.MemoryManager // Memory manager instance for handling subscriptions and transactions
}

// NewParser creates a new instance of Parser with an initialized MemoryManager.
func NewParser() *Parser {
	return &Parser{
		currentBlock:  0,
		memoryManager: memory.NewMemoryManager(),
	}
}

// GetCurrentBlock returns the last parsed Ethereum block number.
func (p *Parser) GetCurrentBlock() int {
	return p.currentBlock
}

// Subscribe adds an address to the subscription list for transaction tracking.
func (p *Parser) Subscribe(address string) bool {
	return p.memoryManager.Subscribe(address)
}

// GetTransactions retrieves all transactions for a subscribed address.
func (p *Parser) GetTransactions(address string) []types.Transaction {
	return p.memoryManager.GetTransactions(address)
}

// UpdateBlock retrieves the latest block number from the Ethereum JSON-RPC endpoint
// and updates the current block in the parser.
func (p *Parser) UpdateBlock() error {
	// Send request to Ethereum RPC to fetch current block number
	reqBody, err := json.Marshal(map[string]interface{}{
		"jsonrpc": config.JSONRPCVersion,
		"method":  config.BlockMethod,
		"params":  []string{},
		"id":      83,
	})

	if err != nil {
		log.Printf("Failed to marshal JSON request: %v", err)
		return err
	}

	response, err := http.Post(config.EthereumRPCURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Printf("Failed to fetch block number: %v", err)
		return err
	}
	defer response.Body.Close()

	var rpcResp types.RPCResponse
	if err := json.NewDecoder(response.Body).Decode(&rpcResp); err != nil {
		log.Printf("Failed to parse JSON response: %v", err)
		return err
	}

	// Log the response to inspect the result
	log.Printf("RPC Response: %+v", rpcResp)

	// Parse the result from hex to int
	blockNumber, err := strconv.ParseInt(rpcResp.Result, 0, 64)
	if err != nil {
		log.Printf("Failed to parse block number: %v", err)
		return err
	}

	// Update the current block
	p.currentBlock = int(blockNumber)
	log.Printf("Updated block to %d", p.currentBlock)

	// Fetch transactions for the updated block
	return p.FetchTransactionsForBlock(int(blockNumber))
}

// FetchTransactionsForBlock fetches all transactions for a given block number
// and filters them for subscribed addresses.
func (p *Parser) FetchTransactionsForBlock(blockNumber int) error {
	hexBlockNumber := fmt.Sprintf("0x%x", blockNumber)

	requestPayload := types.JSONRPCRequest{
		Jsonrpc: config.JSONRPCVersion,
		Method:  config.BlockByNumberMethod,
		Params:  []interface{}{hexBlockNumber, true},
		ID:      1,
	}

	requestBody, err := json.Marshal(requestPayload)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	resp, err := http.Post(config.EthereumRPCURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var response types.JSONRPCResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Loop through transactions and store them for subscribed addresses
	for _, tx := range response.Result.Transactions {
		if p.memoryManager.IsSubscribed(tx.To) {
			p.memoryManager.StoreTransaction(tx.To, tx)
		}
		if p.memoryManager.IsSubscribed(tx.From) {
			p.memoryManager.StoreTransaction(tx.From, tx)
		}
	}

	log.Printf("Fetched %d transactions for block %d", len(response.Result.Transactions), blockNumber)
	return nil
}

// StartBlockUpdater starts a background task that periodically updates the current block.
func (p *Parser) StartBlockUpdater(interval time.Duration) {
	go func() {
		for {
			err := p.UpdateBlock()
			if err != nil {
				log.Printf("Error updating block: %v", err)
			}
			time.Sleep(interval)
		}
	}()
}
