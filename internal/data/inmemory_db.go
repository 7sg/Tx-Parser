package data

import "parser/internal/model"

type InMemoryDB struct {
	subscribers         map[string]bool
	currentBlockNumber  int64
	addressTransactions map[string]map[string]*model.Transaction
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		subscribers:         make(map[string]bool),
		currentBlockNumber:  0,
		addressTransactions: make(map[string]map[string]*model.Transaction),
	}
}

func (db *InMemoryDB) AddSubscriber(address string) error {
	db.subscribers[address] = true
	return nil
}

func (db *InMemoryDB) IsSubscriber(address string) bool {
	_, ok := db.subscribers[address]
	return ok
}

func (db *InMemoryDB) SaveTransaction(address, transactionHash string, transaction *model.Transaction) error {
	if _, ok := db.addressTransactions[address]; !ok {
		db.addressTransactions[address] = make(map[string]*model.Transaction)
	}
	db.addressTransactions[address][transactionHash] = transaction
	return nil
}

func (db *InMemoryDB) GetTransactions(address string) (map[string]*model.Transaction, error) {
	return db.addressTransactions[address], nil
}

func (db *InMemoryDB) GetCurrentBlockNumber() (int64, error) {
	return db.currentBlockNumber, nil
}

func (db *InMemoryDB) SaveCurrentBlockNumber(blockNumber int64) error {
	db.currentBlockNumber = blockNumber
	return nil
}
