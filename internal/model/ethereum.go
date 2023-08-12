package model

type Request struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

type GetBlockByNumberResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  struct {
		Difficulty       string         `json:"difficulty"`
		ExtraData        string         `json:"extraData"`
		GasLimit         string         `json:"gasLimit"`
		GasUsed          string         `json:"gasUsed"`
		Hash             string         `json:"hash"`
		LogsBloom        string         `json:"logsBloom"`
		Miner            string         `json:"miner"`
		MixHash          string         `json:"mixHash"`
		Nonce            string         `json:"nonce"`
		Number           string         `json:"number"`
		ParentHash       string         `json:"parentHash"`
		ReceiptsRoot     string         `json:"receiptsRoot"`
		Sha3Uncles       string         `json:"sha3Uncles"`
		Size             string         `json:"size"`
		StateRoot        string         `json:"stateRoot"`
		Timestamp        string         `json:"timestamp"`
		TotalDifficulty  string         `json:"totalDifficulty"`
		Transactions     []*Transaction `json:"transactions"`
		TransactionsRoot string         `json:"transactionsRoot"`
		Uncles           []interface{}  `json:"uncles"`
	} `json:"result"`

	Error *Error `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GetBlockNumberResponse struct {
	Id      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`

	Error *Error `json:"error"`
}
