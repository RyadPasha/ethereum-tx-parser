// Ethereum Transaction Parser
//
// Created by: Mohamed Riyad
// Date: 2024-11-14
// License: MIT
//

package memory

import (
	"sync"

	"ethereum-tx-parser/types"
)

// MemoryManager manages the subscription list and transactions for subscribed addresses.
type MemoryManager struct {
	subscribed   map[string]bool                // Set of subscribed addresses
	transactions map[string][]types.Transaction // Map of transactions by address
	mutex        sync.Mutex                     // Mutex for concurrent access
}

// NewMemoryManager creates a new instance of MemoryManager.
func NewMemoryManager() *MemoryManager {
	return &MemoryManager{
		subscribed:   make(map[string]bool),
		transactions: make(map[string][]types.Transaction),
	}
}

// Subscribe adds an address to the subscription list and returns true if successful.
// If the address is already subscribed, it returns false.
func (m *MemoryManager) Subscribe(address string) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if _, exists := m.subscribed[address]; exists {
		return false
	}
	m.subscribed[address] = true
	return true
}

// IsSubscribed checks if a given address is subscribed.
func (m *MemoryManager) IsSubscribed(address string) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.subscribed[address]
}

// StoreTransaction stores a transaction for a given address.
func (m *MemoryManager) StoreTransaction(address string, transaction types.Transaction) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.transactions[address] = append(m.transactions[address], transaction)
}

// GetTransactions retrieves all transactions for a subscribed address.
// If the address is not subscribed, it returns an empty list.
func (m *MemoryManager) GetTransactions(address string) []types.Transaction {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.transactions[address]
}
