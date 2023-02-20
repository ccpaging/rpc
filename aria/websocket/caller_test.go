package wscaller

import (
	"context"
	"testing"
	"time"

	"github.com/ccpaging/rpc/aria"
)

func TestWebsocketCaller(t *testing.T) {
	time.Sleep(time.Second)
	c, err := NewWebsocketCaller(context.Background(), "ws://localhost:6800/jsonrpc", time.Second, &DummyNotifier{})
	if err != nil {
		t.Fatal(err.Error())
	}
	defer c.Close()

	var info aria.VersionInfo
	if err := c.Call("aria2.getVersion", []interface{}{}, &info); err != nil {
		t.Error(err.Error())
	} else {
		println(info.Version)
	}
}
