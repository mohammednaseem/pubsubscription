package http

import (
	"time"

	"github.com/gcp-iot/model"
)

type registryService struct {
	contextTimeout time.Duration
	baseUrl        string
}

func NewIoTService(timeout time.Duration, url string) model.IHTTPEndpointInvoker {
	return &registryService{
		contextTimeout: timeout,
		baseUrl:        url,
	}
}
