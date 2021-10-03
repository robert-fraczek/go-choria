package inter

import (
	"context"
	"net/url"
)

// RequestSignerConfig configures RequestSigner
type RequestSignerConfig interface {
	RemoteSignerURL() (*url.URL, error)
	RemoteSignerToken() ([]byte, error)
}

// RequestSigner signs request, typically remote signers over HTTP or Choria RPC
type RequestSigner interface {
	// Sign signs request payload
	Sign(ctx context.Context, request []byte, cfg RequestSignerConfig) ([]byte, error)

	// Kind is the name of the provider
	Kind() string
}