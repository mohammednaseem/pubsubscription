package usecase

import (
	"time"

	"github.com/gcp-iot/model"
)

type registryUsecase struct {
	registryService model.IHTTPEndpointInvoker
	contextTimeout  time.Duration
}

func NewIoTUsecase(r model.IHTTPEndpointInvoker, timeout time.Duration) model.IMessageRouter {
	return &registryUsecase{
		registryService: r,
		contextTimeout:  timeout,
	}
}
