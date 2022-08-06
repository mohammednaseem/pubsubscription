package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
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
			log.Info().Msg("Acknowledged Message")
		}

	})
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}
	log.Info().Msg(fmt.Sprintf("Received %d", received))
	return err
}
