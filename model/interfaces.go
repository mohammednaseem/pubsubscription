package model

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type IMessageRouter interface {
	// message parser
	DecodeMessage(ctx context.Context, message *pubsub.Message) (Response, error)
}

type IHTTPEndpointInvoker interface {
	HttpPost(ctx context.Context, entity string, body string) (Response, error)
	HttpPatch(ctx context.Context, entity string, body string) (Response, error)
	HttpDelete(ctx context.Context, entity string, body string) (Response, error)

	// CreateDevice(ctx context.Context, registry string) (Response, error)
	// UpdateDevice(ctx context.Context, registry string) (Response, error)
	// DeleteDevice(ctx context.Context, registry string) (Response, error)
	// httpPost(ctx context.Context, entity string, body string) (Response, error)
	// httpPatch(ctx context.Context, entity string, body string) (Response, error)
	// httpDelete(ctx context.Context, entity string, body string) (Response, error)
}

type IGRPCEndpointInvoker interface {
	//endpoint invoker
}
