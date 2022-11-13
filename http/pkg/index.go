package pkg

import (
	"net/http"
	"time"
)

// HTTP ...
type HTTP struct {
	Client *http.Client
}

// HTTPConfig ...
type HTTPConfig struct {
	Timeout             int // In milliseconds, If requested resource is not responded within time, then HTTP connection will be canceled with net/http: request canceled (Client.Timeout exceeded ...) error. (Not recommended more than 10 seconds)
	MaxIdleConns        int // This is the connection pool size. This is the max connection that can be open. (Default value is 100 connections)
	IdleConnTimeout     int // In milliseconds, IdleConnTimeout is the maximum amount of time an idle (keep-alive) connection will remain idle before closing itself.(Zero means no limit)
	MaxIdleConnsPerHost int // It is the number of connection can be allowed to open per host basic.
	TLSHandshakeTimeout int // In milliseconds, TLSHandshakeTimeout specifies the maximum amount of time waiting to wait for a TLS handshake. (Zero means no timeout)
}

// New ...
func New(config *HTTPConfig) (httpClient *HTTP) {
	// Checks
	if config == nil {
		// Gets default http client
		httpClient := getDefaultHTTPClient()

		// Returns
		return &HTTP{
			Client: httpClient,
		}
	}

	// Forms http client
	transport := &http.Transport{
		MaxIdleConns:        config.MaxIdleConns,
		IdleConnTimeout:     time.Duration(config.IdleConnTimeout) * time.Millisecond,
		MaxIdleConnsPerHost: config.MaxIdleConnsPerHost,
		TLSHandshakeTimeout: time.Duration(config.TLSHandshakeTimeout) * time.Millisecond,
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(config.Timeout) * time.Millisecond,
	}

	// Returns
	return &HTTP{
		Client: client,
	}
}

func getDefaultHTTPClient() (httpClient *http.Client) {
	// Forms transport
	transport := &http.Transport{
		MaxIdleConns:        100,
		IdleConnTimeout:     time.Duration(1 * time.Minute),
		MaxIdleConnsPerHost: 100,
		TLSHandshakeTimeout: time.Duration(100 * time.Millisecond),
	}

	// Returns
	return &http.Client{
		Transport: transport,
		Timeout:   time.Duration(200 * time.Millisecond),
	}
}
