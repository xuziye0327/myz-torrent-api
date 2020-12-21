package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

const testProtal = 34667

func TestJSONRPC(t *testing.T) {
	startUpLocalServer()
	rpc := NewRPCClient(fmt.Sprintf("http://localhost:%v/rpctest", testProtal))
	t.Run("get", func(t *testing.T) {
		if resp, err := rpc.Call("get"); err != nil {
			t.Error(err)
		} else {
			if resp.Result != "get" {
				t.Errorf("Restlt is not 'get'")
			}
		}
	})
}

func startUpLocalServer() {
	http.HandleFunc("/rpctest", func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		var rpcReq RPCRequest
		if err := decoder.Decode(&rpcReq); err != nil {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(200)
			bs, _ := json.Marshal(RPCResponse{
				Result: rpcReq.Method,
			})
			w.Write(bs)
		}
	})

	go http.ListenAndServe(fmt.Sprintf(":%v", testProtal), nil)
}
