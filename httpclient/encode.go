package httpjson2

import (
	"bytes"
	"encoding/json"
	"sync/atomic"
	"time"
)

// clientRequest represents a JSON-RPC request sent by a client.
type clientRequest struct {
	// JSON-RPC protocol.
	Version string `json:"jsonrpc"`

	// A String containing the name of the method to be invoked.
	Method string `json:"method"`

	// Object to pass as request parameter to the method.
	Params interface{} `json:"params"`

	// The request id. This can be of any type. It is used to match the
	// response with the request that it is replying to.
	Id uint64 `json:"id"`
}

var reqId = func() func() uint64 {
	var id = uint64(time.Now().UnixNano())
	return func() uint64 {
		return atomic.AddUint64(&id, 1)
	}
}()

func encodeRequest(method string, args interface{}) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	c := &clientRequest{
		Version: "2.0",
		Method:  method,
		Params:  args,
		Id:      reqId(),
	}
	if err := json.NewEncoder(&buf).Encode(c); err != nil {
		return nil, err
	}
	return &buf, nil
}
