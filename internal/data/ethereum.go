package data

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"parser/internal/model"
)

type ethereum struct {
	url string
}

func NewEthereum() *ethereum {
	return &ethereum{
		url: "https://cloudflare-eth.com",
	}
}

func (e *ethereum) GetBlockByNumber(blockNumber string, needTransactionDetails bool) (*model.GetBlockByNumberResponse, error) {
	request := model.Request{
		Jsonrpc: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{blockNumber, needTransactionDetails},
		Id:      1,
	}
	marshal, err := json.Marshal(request)
	if err != nil {
		log.Printf("error while marshaling request: %v", err)
		return nil, err
	}
	resp, err := http.Post(e.url, "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		log.Printf("error while sending request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error while reading response body: %v", err)
		return nil, err
	}

	var response *model.GetBlockByNumberResponse
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		log.Printf("error while unmarshaling response: %v", err)
		return nil, err
	}

	return response, nil
}

func (e *ethereum) GetCurrentBlockNumber() (*model.GetBlockNumberResponse, error) {
	request := model.Request{
		Jsonrpc: "2.0",
		Method:  "eth_blockNumber",
		Id:      1,
	}
	marshal, err := json.Marshal(request)
	if err != nil {
		log.Printf("error while marshaling request: %v", err)
		return nil, err
	}
	resp, err := http.Post(e.url, "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		log.Printf("error while sending request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error while reading response body: %v", err)
		return nil, err
	}

	var response *model.GetBlockNumberResponse
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		log.Printf("error while unmarshaling response: %v", err)
		return nil, err
	}

	return response, nil
}
