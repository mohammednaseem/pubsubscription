package model

import "context"

type IMessageRouter interface {
	// message parser
	DecodeMessage(ctx context.Context, message string) (Response, error)
}

type IHTTPEndpointInvoker interface {
	FormHTTPRequestAndInvokeEndpoint(ctx context.Context, message string) (Response, error)
}

type IGRPCEndpointInvoker interface {
	//endpoint invoker
}
