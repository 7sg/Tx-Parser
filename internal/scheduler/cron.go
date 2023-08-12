package scheduler

import (
	"log"
	"parser/internal/business"
	"time"
)

type Scheduler struct {
	interval time.Duration
	parser   *business.Parser
}

func NewScheduler(parser *business.Parser) *Scheduler {
	return &Scheduler{
		interval: 1 * time.Second,
		parser:   parser,
	}
}

func (s *Scheduler) Run() {
	ticker := time.NewTicker(s.interval)
	for range ticker.C {
		err := s.parser.SyncTransactions()
		if err != nil {
			log.Printf("error while syncing transactions: %v", err)
		}
	}
}
