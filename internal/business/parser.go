package business

import (
	"fmt"
	"log"
	"parser/internal/model"
	"parser/internal/util"
)

type database interface {
	AddSubscriber(address string) error
	IsSubscriber(address string) bool
	SaveTransaction(address, transactionHash string, transaction *model.Transaction) error
	GetTransactions(address string) (map[string]*model.Transaction, error)
	GetCurrentBlockNumber() (int64, error)
	SaveCurrentBlockNumber(blockNumber int64) error
}

type ethereum interface {
	GetBlockByNumber(blockNumber string, needTransactionDetails bool) (*model.GetBlockByNumberResponse, error)
	GetCurrentBlockNumber() (*model.GetBlockNumberResponse, error)
}

type Parser struct {
	db  database
	eth ethereum
}

func NewParser(db database, eth ethereum) *Parser {
	return &Parser{
		db:  db,
		eth: eth,
	}
}

func (p *Parser) InitCurrentBlockNumber() error {
	currentBlockNumber, err := p.getCurrentBlockNumberHelper()
	if err != nil {
		return err
	}
	return p.db.SaveCurrentBlockNumber(currentBlockNumber)
}

func (p *Parser) SyncTransactions() error {
	currentBlockNumber, err := p.getCurrentBlockNumberHelper()
	if err != nil {
		return err
	}

	inProgressBlockNumber, err := p.db.GetCurrentBlockNumber()
	if err != nil {
		return err
	}

	for inProgressBlockNumber <= currentBlockNumber {
		log.Printf("sync for block %d", inProgressBlockNumber)
		blockNumberHex := util.IntegerToHex(inProgressBlockNumber)
		block, err := p.eth.GetBlockByNumber(blockNumberHex, true)
		if err != nil {
			return err
		}
		for _, transaction := range block.Result.Transactions {
			if p.db.IsSubscriber(transaction.From) {
				log.Printf("sync for subscriber %s", transaction.From)
				err = p.db.SaveTransaction(transaction.From, transaction.Hash, transaction)
				if err != nil {
					return err
				}
			}
			if p.db.IsSubscriber(transaction.To) {
				log.Printf("sync for subscriber %s", transaction.To)
				err = p.db.SaveTransaction(transaction.To, transaction.Hash, transaction)
				if err != nil {
					return err
				}
			}
		}
		inProgressBlockNumber++
		p.db.SaveCurrentBlockNumber(inProgressBlockNumber)
	}
	return nil
}

func (p *Parser) GetCurrentBlock() (int64, error) {
	inProgressBlockNumber, err := p.db.GetCurrentBlockNumber()
	if err != nil {
		return 0, err
	}
	return inProgressBlockNumber, nil
}

func (p *Parser) Subscribe(address string) error {
	return p.db.AddSubscriber(address)
}

func (p *Parser) GetTransactions(address string) ([]*model.Transaction, error) {
	transactions, err := p.db.GetTransactions(address)
	if err != nil {
		return nil, err
	}
	transactionsList := make([]*model.Transaction, 0, len(transactions))
	for _, transaction := range transactions {
		transactionsList = append(transactionsList, transaction)
	}
	return transactionsList, nil
}

func (p *Parser) getCurrentBlockNumberHelper() (int64, error) {
	response, err := p.eth.GetCurrentBlockNumber()
	if err != nil {
		return 0, err
	}
	if response.Error != nil {
		return 0, fmt.Errorf(response.Error.Message)
	}
	currentBlockNumber, err := util.HexToInteger(response.Result)
	if err != nil {
		return 0, err
	}
	return currentBlockNumber, nil
}
