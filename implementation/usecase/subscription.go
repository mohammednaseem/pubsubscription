package usecase

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/gcp-iot/model"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
)

type Publish struct {
	Operation string `json:"operation" validate:"required"`
	Entity    string `json:"entity" validate:"required"`
}

// createRegistry creates a IoT Core device registry associated with a PubSub topic
func (i *registryUsecase) DecodeMessage(ctx context.Context, msg *pubsub.Message) (model.Response, error) {
	// c, cancel := context.WithTimeout(ctx, i.contextTimeout)
	// defer cancel()

	// dr, err := i.registryService.CreateRegistry(c, msg)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return dr, err

	// }
	operation := gjson.Get(string(msg.Data), "operation").String()
	entity := gjson.Get(string(msg.Data), "entity").String()
	data := gjson.Get(string(msg.Data), "data").String()
	log.Info().Msg(operation + entity)
	var err error
	switch entity {
	case "Registry":
		entity := "registry"
		switch operation {
		case "CREATE":
			i.registryService.HttpPost(ctx, entity, data)
		case "UPDATE":
			i.registryService.HttpPatch(ctx, entity, data)
		case "DELETE":
			i.registryService.HttpDelete(ctx, entity, data)
		default:
			log.Error().Msg("Unknown Operation in Published Message Registry ")

		}
	case "Device":
		entity := "device"
		switch operation {
		case "CREATE":
			i.registryService.HttpPost(ctx, entity, data)
		case "UPDATE":
			i.registryService.HttpPatch(ctx, entity, data)
		case "DELETE":
			i.registryService.HttpDelete(ctx, entity, data)
		default:
			log.Error().Msg("Unknown Operation in Published Message Device")

		}
	default:
		log.Error().Msg("Unknown Entity in Published Message")

	}

	var dr model.Response
	return dr, err
}
