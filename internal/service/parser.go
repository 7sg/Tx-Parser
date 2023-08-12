package service

import (
	"encoding/json"
	"io"
	"net/http"
	"parser/internal/business"
)

type ParserService struct {
	parser *business.Parser
}

func NewParserService(parser *business.Parser) *ParserService {
	return &ParserService{
		parser: parser,
	}
}

func (ps *ParserService) GetCurrentBlock(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	currentBlock, err := ps.parser.GetCurrentBlock()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := GetCurrentBlockResponse{
		CurrentBlock: currentBlock,
	}
	json.NewEncoder(w).Encode(response)
}

func (ps *ParserService) Subscribe(w http.ResponseWriter, r *http.Request) {
	response := SubscribeResponse{}
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	var subscribeRequest SubscribeRequest
	err = json.Unmarshal(body, &subscribeRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	err = ps.parser.Subscribe(subscribeRequest.Address)
	if err != nil {
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.Success = true
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (ps *ParserService) GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	address := r.URL.Query().Get("address")
	transactions, err := ps.parser.GetTransactions(address)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

type SubscribeRequest struct {
	Address string `json:"address"`
}

type SubscribeResponse struct {
	Success bool `json:"success"`
}

type GetCurrentBlockResponse struct {
	CurrentBlock int64 `json:"current_block"`
}
