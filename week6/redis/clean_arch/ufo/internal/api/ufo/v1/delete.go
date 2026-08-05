package v1

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	ufo_v1 "github.com/mbakhodurov/examples2/week_6/redis/clean_arch/shared/pkg/proto/ufo/v1"
	errs "github.com/mbakhodurov/examples2/week_6/redis/clean_arch/ufo/internal/errors"
)

func (a *api) Delete(ctx context.Context, req *ufo_v1.DeleteRequest) (*ufo_v1.DeleteResponse, error) {
	err := a.ufoService.Delete(ctx, req.GetUuid())
	if err != nil {
		if errors.Is(err, errs.ErrSightingNotFound) {
			return nil, status.Errorf(codes.NotFound, "наблюдение с UUID %s не найдено", req.GetUuid())
		}
		return nil, err
	}

	return &ufo_v1.DeleteResponse{}, nil
}
