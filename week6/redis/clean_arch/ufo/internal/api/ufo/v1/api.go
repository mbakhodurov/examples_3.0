package v1

import ufo_v1 "github.com/mbakhodurov/examples2/week_6/redis/clean_arch/shared/pkg/proto/ufo/v1"

// api реализует gRPC-обработчики сервиса наблюдений НЛО
type api struct {
	ufo_v1.UnimplementedUFOServiceServer

	ufoService UFOService
}

// New создаёт API-обработчик наблюдений НЛО
func New(ufoService UFOService) *api {
	return &api{
		ufoService: ufoService,
	}
}
