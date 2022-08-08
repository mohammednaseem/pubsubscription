package http

import (
	"time"

	"github.com/gcp-iot/model"
)

type subService struct {
	contextTimeout time.Duration
}

func NewSubService(timeout time.Duration) model.IHTTPEndpointInvoker {
	return &subService{
		contextTimeout: timeout,
	}
}
