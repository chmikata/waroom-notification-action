package incident

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

type Requester interface {
	ExecHttpReq(req *http.Request) ([]byte, error)
}

type HttpRequest struct {
	client *http.Client
}

func NewHttpRequest() *HttpRequest {
	return &HttpRequest{
		client: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout: 10 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout:   10 * time.Second,
				ResponseHeaderTimeout: 10 * time.Second,
				IdleConnTimeout:       10 * time.Second,
				MaxIdleConns:          100,
				MaxConnsPerHost:       100,
				MaxIdleConnsPerHost:   100,
				TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}
