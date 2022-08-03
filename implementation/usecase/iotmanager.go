package usecase

import (
	"time"

	"github.com/gcp-iot/model"
)

type registryUsecase struct {
	registryService model.IHTTPEndpointInvoker
	contextTimeout  time.Duration
	baseUrl         string
}

func NewIoTUsecase(r model.IHTTPEndpointInvoker, timeout time.Duration, url string) model.IMessageRouter {
	return &registryUsecase{
		registryService: r,
		contextTimeout:  timeout,
		baseUrl:         url,
	}
}
