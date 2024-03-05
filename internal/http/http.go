package http

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

var httpClient *http.Client

func newClient(roundTripper http.RoundTripper) *http.Client {
	return &http.Client{
		Transport: otelhttp.NewTransport(roundTripper),
	}
}

func NewHttpClient(roundTripper *http.RoundTripper) *http.Client {
	if roundTripper != nil {
		return newClient(*roundTripper)
	}

	if httpClient == nil {
		httpClient = newClient(http.DefaultTransport)
	}
	return httpClient
}
