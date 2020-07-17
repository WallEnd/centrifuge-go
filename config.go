package centrifuge

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

const (
	// DefaultHandshakeTimeout ...
	DefaultHandshakeTimeout = time.Second
	// DefaultReadTimeout ...
	DefaultReadTimeout = 5 * time.Second
	// DefaultWriteTimeout ...
	DefaultWriteTimeout = time.Second
	// DefaultPingInterval ...
	DefaultPingInterval = 25 * time.Second
	// DefaultPrivateChannelPrefix ...
	DefaultPrivateChannelPrefix = "$"
	// NumReconnect is maximum number of reconnect attempts, 0 means reconnect forever.
	DefaultBackoffNumReconnect = 0
	// MinMilliseconds is a minimum value of the reconnect interval.
	DefaultBackoffMinMilliseconds = 100
	// MaxMilliseconds is a maximum value of the reconnect interval.
	DefaultBackoffMaxMilliseconds = 20 * 1000
	// Factor is the multiplying factor for each increment step.
	DefaultBackoffFactor = 2
	// Jitter eases contention by randomizing backoff steps.
	DefaultBackoffJitter = true
)

// WsConfig contains various client options.
type WsConfig struct {
	// NetDialContext specifies the dial function for creating TCP connections. If
	// NetDialContext is nil, net.DialContext is used.
	NetDialContext func(ctx context.Context, network, addr string) (net.Conn, error)

	// PrivateChannelPrefix is private channel prefix.
	PrivateChannelPrefix string
	// ReadTimeout is how long to wait read operations to complete.
	ReadTimeout time.Duration
	// WriteTimeout is Websocket write timeout.
	WriteTimeout time.Duration
	// PingInterval is how often to send ping commands to server.
	PingInterval time.Duration
	// HandshakeTimeout specifies the duration for the handshake to complete.
	HandshakeTimeout time.Duration
	// TLSConfig specifies the TLS configuration to use with tls.Client.
	// If nil, the default configuration is used.
	TLSConfig *tls.Config
	// EnableCompression specifies if the client should attempt to negotiate
	// per message compression (RFC 7692). Setting this value to true does not
	// guarantee that compression will be supported. Currently only "no context
	// takeover" modes are supported.
	EnableCompression bool
	// CookieJar specifies the cookie jar.
	// If CookieJar is nil, cookies are not sent in requests and ignored
	// in responses.
	CookieJar http.CookieJar
	// Header specifies custom HTTP Header to send.
	Header http.Header
}

type BackoffReconnectConfig struct {
	// NumReconnect is maximum number of reconnect attempts, 0 means reconnect forever.
	NumReconnect int
	// Factor is the multiplying factor for each increment step.
	Factor float64
	// Jitter eases contention by randomizing backoff steps.
	Jitter bool
	// MinMilliseconds is a minimum value of the reconnect interval.
	MinMilliseconds int
	// MaxMilliseconds is a maximum value of the reconnect interval.
	MaxMilliseconds int
}

type Config struct {
	BackoffConfig BackoffReconnectConfig
	WsConfig      WsConfig
}

// DefaultConfig returns WsConfig with default options.
func DefaultConfig() Config {
	return Config{
		BackoffReconnectConfig{
			NumReconnect:    DefaultBackoffNumReconnect,
			Factor:          DefaultBackoffFactor,
			Jitter:          DefaultBackoffJitter,
			MinMilliseconds: DefaultBackoffMinMilliseconds,
			MaxMilliseconds: DefaultBackoffMaxMilliseconds,
		},
		WsConfig{
			PingInterval:         DefaultPingInterval,
			ReadTimeout:          DefaultReadTimeout,
			WriteTimeout:         DefaultWriteTimeout,
			HandshakeTimeout:     DefaultHandshakeTimeout,
			PrivateChannelPrefix: DefaultPrivateChannelPrefix,
			Header:               http.Header{},
		},
	}
}
