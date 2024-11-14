// Ethereum Transaction Parser
//
// Created by: Mohamed Riyad
// Date: 2024-11-14
// License: MIT
//

package types

// Transaction represents a transaction on the Ethereum blockchain.
type Transaction struct {
	Hash                 string `json:"hash"`
	Nonce                string `json:"nonce"`
	BlockHash            string `json:"blockHash"`
	BlockNumber          string `json:"blockNumber"`
	TransactionIndex     string `json:"transactionIndex"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	Value                string `json:"value"`
	GasPrice             string `json:"gasPrice"`
	Gas                  string `json:"gas"`
	MaxFeePerGas         string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
	Input                string `json:"input"`
	R                    string `json:"r"`
	S                    string `json:"s"`
	V                    string `json:"v"`
	YParity              string `json:"yParity"`
	ChainId              string `json:"chainId"`
	Type                 string `json:"type"`
}
type ResultDetail struct {
	Size         string        `json:"size"`
	Transactions []Transaction `json:"transactions"`
}
type JSONRPCResponse struct {
	Jsonrpc string       `json:"jsonrpc"`
	ID      int          `json:"id"`
	Result  ResultDetail `json:"result"`
}

// RPCResponse represents the response format for the JSON-RPC call to get the current block number.
type RPCResponse struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Result  string `json:"result"` // Hex value of the block number
}

// JSONRPCRequest represents the structure of our JSON-RPC request
type JSONRPCRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}
