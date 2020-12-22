package common

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const jsonrpcVersion = "2.0"

// RPCRequest is represented by sending a Request object to a Server.
// JSON-RPC 2.0 Specification
type RPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      string      `json:"id"`
}

// RPCResponse is expressed as a single JSON Object, the Server MUST reply with a Response.
type RPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
	ID      string      `json:"id"`
}

// RPCError is a Object that the Response Object MUST contain the error member with a value
type RPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// RPCClient is a json rpc 2.0 client
type RPCClient struct {
	url        string
	httpClient *http.Client
}

// NewRPCClient creates a rpc client
func NewRPCClient(url string) *RPCClient {
	httpCli := &http.Client{
		Timeout: time.Second * 10,
	}

	return &RPCClient{
		url:        url,
		httpClient: httpCli,
	}
}

// Call a jsonrpc method
func (cli *RPCClient) Call(method string, params ...interface{}) (*RPCResponse, error) {
	rpcReq := &RPCRequest{
		JSONRPC: jsonrpcVersion,
		Method:  method,
	}
	if len(params) != 0 {
		rpcReq.Params = params
	}

	rpcJSON, err := json.Marshal(rpcReq)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", cli.url, bytes.NewReader(rpcJSON))
	if err != nil {
		return nil, err
	}

	httpResp, err := cli.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	var rpcResp *RPCResponse
	decoder := json.NewDecoder(httpResp.Body)
	err = decoder.Decode(&rpcResp)
	log.Printf("jsonrpc result: %s, err: %v", JsonToString(rpcResp), err)

	return rpcResp, err
}
