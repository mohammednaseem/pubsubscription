package http

import (
	"time"

	"github.com/gcp-iot/model"
)

type registryService struct {
	contextTimeout time.Duration
}

func NewIoTService(timeout time.Duration) model.IHTTPEndpointInvoker {
	return &registryService{
		contextTimeout: timeout,
	}
}
