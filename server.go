package websocket

import (
	"errors"
	"strings"
)

var (
	keyGUID = []byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11")
	// ErrBadRequestMethod bad request method
	ErrBadRequestMethod = errors.New("bad method")
	// ErrNotWebSocket not websocket protocal
	ErrNotWebSocket = errors.New("not websocket protocol")
	// ErrBadWebSocketVersion bad websocket version
	ErrBadWebSocketVersion = errors.New("missing or bad WebSocket Version")
	// ErrChallengeResponse mismatch challenge response
	ErrChallengeResponse = errors.New("mismatch challenge/response")
)

func Upgrade(req *Request) (conn *Conn, err error) {
	if req.Method != "GET" {
		return nil, ErrBadRequestMethod
	}
	if req.Header.Get("Sec-Websocket-Version") != "13" {
		return nil, ErrBadWebSocketVersion
	}
	if strings.ToLower(req.Header.Get("Upgrade")) != "websocket" {
		return nil, ErrNotWebSocket
	}
	if !strings.Contains(strings.ToLower(req.Header.Get("Connection")), "upgrade") {
		return nil, ErrNotWebSocket
	}
	challengeKey := req.Header.Get("Sec-Websocket-Key")
	if challengeKey == "" {
		return nil, ErrChallengeResponse
	}

}
