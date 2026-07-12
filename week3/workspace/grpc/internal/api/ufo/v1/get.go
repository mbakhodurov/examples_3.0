package v1

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mbakhodurov/examples2/week_3/workspace/grpc/internal/api/converter"
	errs "github.com/mbakhodurov/examples2/week_3/workspace/grpc/internal/errors"
	ufov1 "github.com/mbakhodurov/examples2/week_3/workspace/shared/pkg/proto/ufo/v1"
)

func (a *api) Get(ctx context.Context, req *ufov1.GetRequest) (*ufov1.GetResponse, error) {
	sighting, err := a.ufoRepository.Get(ctx, req.GetUuid())
	if err != nil {
		if errors.Is(err, errs.ErrSightingNotFound) {
			return nil, status.Errorf(codes.NotFound, "наблюдение с UUID %s не найдено", req.GetUuid())
		}
		return nil, err
	}

	return &ufov1.GetResponse{
		Sighting: converter.SightingToDTO(sighting),
	}, nil
}
