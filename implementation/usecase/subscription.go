package usecase

import (
	"context"
	"errors"

	"cloud.google.com/go/pubsub"
	"github.com/gcp-iot/model"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
)

// createRegistry creates a IoT Core device registry associated with a PubSub topic
func (i *subUsecase) DecodeMessage(ctx context.Context, msg *pubsub.Message) (model.Response, error) {

	entity := gjson.Get(string(msg.Data), "entity").String()
	data := gjson.Get(string(msg.Data), "data").String()
	path := gjson.Get(string(msg.Data), "path").String()
	method := gjson.Get(string(msg.Data), "operation").String()
	if entity == "" || data == "" || path == "" || method == "" {
		log.Error().Msg("Publish Format Unknown")
		return model.Response{}, errors.New("unknown publish")
	}
	log.Info().Msg(method + entity + path)
	dr, err := i.subService.HttpOperation(ctx, method, entity, data, i.baseUrl+path)
	if err != nil {
		log.Error().Err(err).Msg("")
		return dr, err
	}
	return dr, err
}
