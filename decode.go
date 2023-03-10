package rpc

import (
	"encoding/json"
	"io"
)

// clientResponse represents a JSON-RPC response returned to a client.
type ClientResponse struct {
	Version string           `json:"jsonrpc"`
	Result  *json.RawMessage `json:"result"`
	Error   *json.RawMessage `json:"error"`
	Id      *uint64          `json:"id"`
}

func (c ClientResponse) Decode(reply interface{}) error {
	if c.Error != nil {
		jsonErr := &Error{}
		if err := json.Unmarshal(*c.Error, jsonErr); err != nil {
			return &Error{
				Code:    E_SERVER,
				Message: string(*c.Error),
			}
		}
		return jsonErr
	}

	if c.Result == nil {
		return ErrNullResult
	}

	return json.Unmarshal(*c.Result, reply)
}

func decodeResponse(r io.Reader, reply interface{}) error {
	var c ClientResponse
	if err := json.NewDecoder(r).Decode(&c); err != nil {
		return err
	}
	return c.Decode(reply)
}
