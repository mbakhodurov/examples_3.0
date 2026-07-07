package v1

import ufo_v1 "github.com/mbakhodurov/examples2/week_2/layers/pkg/proto/ufo/v1"

type api struct {
	ufo_v1.UnimplementedUFOServiceServer

	ufoService UFOService
}

func NewApi(ufoService UFOService) *api {
	return &api{
		ufoService: ufoService,
	}
}
