// Ethereum Transaction Parser
//
// Created by: Mohamed Riyad
// Date: 2024-11-14
// License: MIT
//

package config

// General application constants
const (
	DefaultAppPort = "8080"                                // Port for HTTP server
	EthereumRPCURL = "https://ethereum-rpc.publicnode.com" // Ethereum RPC endpoint
)

// JSON-RPC related constants
const (
	JSONRPCVersion      = "2.0"
	BlockMethod         = "eth_blockNumber"
	BlockByNumberMethod = "eth_getBlockByNumber"
)
