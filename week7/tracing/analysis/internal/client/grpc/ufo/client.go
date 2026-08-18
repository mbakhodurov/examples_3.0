package ufo

import (
	"context"

	"github.com/mbakhodurov/homeworks2/week7/tracing/analysis/internal/model"
	ufov1 "github.com/mbakhodurov/homeworks2/week7/tracing/shared/pkg/proto/ufo/v1"
)

type client struct {
	client ufov1.UFOServiceClient
}

// New создаёт новый клиент для UFO сервиса
func New(grpcClient ufov1.UFOServiceClient) *client {
	return &client{
		client: grpcClient,
	}
}

// GetSighting получает данные о наблюдении по UUID
func (c *client) GetSighting(ctx context.Context, uuid string) (model.Sighting, error) {
	resp, err := c.client.Get(ctx, &ufov1.GetRequest{Uuid: uuid})
	if err != nil {
		return model.Sighting{}, err
	}

	return model.Sighting{
		UUID:            resp.Sighting.Uuid,
		Description:     resp.Sighting.Description,
		Color:           resp.Sighting.Color.GetValue(),
		DurationSeconds: resp.Sighting.DurationSeconds.GetValue(),
	}, nil
}
