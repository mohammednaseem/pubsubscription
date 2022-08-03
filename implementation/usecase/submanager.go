package usecase

import (
	"time"

	"github.com/gcp-iot/model"
)

type subUsecase struct {
	subService     model.IHTTPEndpointInvoker
	contextTimeout time.Duration
	baseUrl        string
}

func NewSubUsecase(r model.IHTTPEndpointInvoker, timeout time.Duration, url string) model.IMessageRouter {
	return &subUsecase{
		subService:     r,
		contextTimeout: timeout,
		baseUrl:        url,
	}
}
