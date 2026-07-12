package v1

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errs "github.com/mbakhodurov/examples2/week_3/workspace/grpc/internal/errors"
	ufov1 "github.com/mbakhodurov/examples2/week_3/workspace/shared/pkg/proto/ufo/v1"
)

func (a *api) Delete(ctx context.Context, req *ufov1.DeleteRequest) (*ufov1.DeleteResponse, error) {
	err := a.ufoRepository.Delete(ctx, req.GetUuid())
	if err != nil {
		if errors.Is(err, errs.ErrSightingNotFound) {
			return nil, status.Errorf(codes.NotFound, "наблюдение с UUID %s не найдено", req.GetUuid())
		}
		return nil, err
	}

	return &ufov1.DeleteResponse{}, nil
}
