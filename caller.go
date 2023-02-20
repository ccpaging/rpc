package rpc

import (
	"net/http"
)

type Caller interface {
	// Call sends a request of rpc to aria2 daemon
	Call(method string, params, reply interface{}) (err error)
	Close() error
}

type HttpCaller struct {
	uri string
	cli *http.Client
}

func NewHTTPCaller(uri string, client *http.Client) *HttpCaller {
	return &HttpCaller{
		uri: uri,
		cli: client,
	}
}

func (h *HttpCaller) Call(method string, params, reply interface{}) (err error) {
	payload, err := encodeRequest(method, params)
	if err != nil {
		return
	}
	r, err := h.cli.Post(h.uri, "application/json", payload)
	if err != nil {
		return
	}
	err = decodeResponse(r.Body, &reply)
	r.Body.Close()
	return
}

func (h *HttpCaller) Close() (err error) {
	return
}
