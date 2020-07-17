package centrifuge

import "errors"

var (
	// ErrTimeout ...
	ErrTimeout = errors.New("timeout")
	// ErrClientClosed ...
	ErrClientClosed = errors.New("client closed")
	// ErrClientDisconnected ...
	ErrClientDisconnected = errors.New("client disconnected")
	// ErrClientExpired ...
	ErrClientExpired = errors.New("client connection expired")
	// ErrReconnectFailed ...
	ErrReconnectFailed = errors.New("reconnect failed")
	// ErrDuplicateSubscription ...
	ErrDuplicateSubscription = errors.New("duplicate subscription")
	//ErrConnectionClosed
	ErrClientDestroyed = errors.New("client destroyed")
)
