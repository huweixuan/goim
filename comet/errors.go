package main

import (
	"errors"
)

var (
	// server
	ErrHandshake = errors.New("handshake failed")
	ErrOperation = errors.New("request operation not valid")
	// codec
	ErrProtoPackLen   = errors.New("default server codec pack length error")
	ErrProtoHeaderLen = errors.New("default server codec header length error")
	// ring
	ErrRingEmpty = errors.New("ring buffer empty")
	ErrRingFull  = errors.New("ring buffer full")
	// timer
	ErrTimerFull   = errors.New("timer full")
	ErrTimerEmpty  = errors.New("timer empty")
	ErrTimerNoItem = errors.New("timer item not exist")
	// channel
	// channel
	ErrPushMsgArg   = errors.New("rpc pushmsg arg error")
	ErrPushMsgsArg  = errors.New("rpc pushmsgs arg error")
	ErrMPushMsgArg  = errors.New("rpc mpushmsg arg error")
	ErrMPushMsgsArg = errors.New("rpc mpushmsgs arg error")
	// rpc
	ErrLogic = errors.New("logic rpc is not available")
)
