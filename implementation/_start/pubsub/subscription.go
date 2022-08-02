package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/rs/zerolog/log"
)

func (r *iSubscriptionHandler) NewMessageHandler(sub *pubsub.Subscription, ctx context.Context) error {
	var received int32
	err := sub.Receive(ctx, func(cctx context.Context, msg *pubsub.Message) {
		//fmt.Println("Got message: " + string(msg.Data))
		received = received + 1
		cctx, cancel := context.WithCancel(cctx)
		_, err := r.MessageRouter.DecodeMessage(ctx, msg)
		if err != nil {
			log.Error().Err(err).Msg("")
			cancel()
		} else {
			msg.Ack()
		}

	})
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	log.Info().Msg(fmt.Sprintf("Received %d", received))
	return err
}
