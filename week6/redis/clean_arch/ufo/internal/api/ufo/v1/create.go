package v1

import (
	"context"

	ufo_v1 "github.com/mbakhodurov/examples2/week_6/redis/clean_arch/shared/pkg/proto/ufo/v1"
	"github.com/mbakhodurov/examples2/week_6/redis/clean_arch/ufo/internal/api/converter"
)

func (a *api) Create(ctx context.Context, req *ufo_v1.CreateRequest) (*ufo_v1.CreateResponse, error) {
	uuid, err := a.ufoService.Create(ctx, converter.CreateRequestToInput(req))
	if err != nil {
		return nil, err
	}

	return &ufo_v1.CreateResponse{
		Uuid: uuid,
	}, nil
}
