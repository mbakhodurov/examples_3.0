package v1

import (
	"context"
	"errors"

	ufov1 "github.com/mbakhodurov/examples2/week_4/di/shared/pkg/proto/ufo/v1"
	errs "github.com/mbakhodurov/examples2/week_4/di/ufo/internal/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *api) Delete(ctx context.Context, req *ufov1.DeleteRequest) (*ufov1.DeleteResponse, error) {
	err := a.ufoService.Delete(ctx, req.GetUuid())
	if err != nil {
		if errors.Is(err, errs.ErrSightingNotFound) {
			return nil, status.Errorf(codes.NotFound, "наблюдение с UUID %s не найдено", req.GetUuid())
		}
		return nil, err
	}

	return &ufov1.DeleteResponse{}, nil
}
