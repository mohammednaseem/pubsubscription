package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/gcp-iot/model"
	"github.com/rs/zerolog/log"
)

type iSubscriptionHandler struct {
	MessageRouter model.IMessageRouter
}

func NewSubHandler(subID string, projectID string, m model.IMessageRouter) {
	SubscriptionHandler := &iSubscriptionHandler{
		MessageRouter: m,
	}
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	defer client.Close()

	sub := client.Subscription(subID)

	// Turn on synchronous mode. This makes the subscriber use the Pull RPC rather than the StreamingPull RPC, which is useful for guaranteeing MaxOutstandingMessages, the max number of messages the client will hold in memory at a time.
	sub.ReceiveSettings.Synchronous = true
	sub.ReceiveSettings.MaxOutstandingMessages = 100

	SubscriptionHandler.NewMessageHandler(sub, ctx)

}
