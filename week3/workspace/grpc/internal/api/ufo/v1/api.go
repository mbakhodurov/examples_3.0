package v1

import ufo_v1 "github.com/mbakhodurov/examples2/week_3/workspace/shared/pkg/proto/ufo/v1"

type api struct {
	ufo_v1.UnimplementedUFOServiceServer

	ufoRepository UFORepository
}

func NewAPI(ufoRepository UFORepository) *api {
	return &api{
		ufoRepository: ufoRepository,
	}
}
