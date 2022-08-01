package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func (r *iSubscriptionHandler) NewMessageHandler(sub *pubsub.Subscription, ctx context.Context) error {
	var received int32
	err := sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		fmt.Println("Got message: " + string(msg.Data))
		received = received + 1
		// call usecase mResponse, err := r.rUsecase.CreateRegistry(ctx, string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(received)
	return nil
}
