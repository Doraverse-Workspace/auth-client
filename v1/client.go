package v1

import (
	"crypto/tls"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	BaseURL string
	TLS     TLSConfig
	IsDebug bool
}

type ClientConfig struct {
	BaseURL string
	TLS     TLSConfig
	IsDebug bool
}

type TLSConfig struct {
	CertFile string
	KeyFile  string
}

var (
	client *Client
)

// NewClient creates a new Client
// returns a Client and an error if the client cannot be created
func NewClient(config ClientConfig) *Client {
	return &Client{
		BaseURL: config.BaseURL,
		TLS:     config.TLS,
		IsDebug: config.IsDebug,
	}
}

// GetClient returns the global Client
// returns a Client and an error if the client cannot be created
func GetClient() *Client {
	return client
}

// NewRequest creates a new Resty client with TLS certificates
// returns a Resty client and an error if the client cannot be created
func (s Client) NewRequest() (*resty.Client, error) {
	client := resty.New()
	crt, err := tls.X509KeyPair([]byte(s.TLS.CertFile), []byte(s.TLS.KeyFile))
	if err != nil {
		Errorf(err, "failed to load client cert", s.IsDebug)
		return nil, fmt.Errorf("failed to load client cert: %w", err)
	}
	client.SetCertificates(crt)

	return client, nil
}
