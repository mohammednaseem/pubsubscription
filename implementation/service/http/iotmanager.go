package http

import (
	"time"

	"github.com/gcp-iot/model"
)

type registryUsecase struct {
	invokerService model.IHTTPEndpointInvoker
	contextTimeout time.Duration
}

func NewIoTUsecase(i model.IHTTPEndpointInvoker, timeout time.Duration) model.IMessageRouter {
	return &registryUsecase{
		invokerService: i,
		contextTimeout: timeout,
	}
}
