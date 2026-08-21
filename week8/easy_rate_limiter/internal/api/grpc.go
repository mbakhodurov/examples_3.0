package api

import (
	"context"
	"log/slog"

	ufo_v1 "github.com/mbakhodurov/examples2/week_8/easy_rate_limiter/pkg/proto/ufo/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateUfo создаёт наблюдение (gRPC)
func (a *api) CreateUfo(_ context.Context, req *ufo_v1.CreateUfoRequest) (*ufo_v1.CreateUfoResponse, error) {
	if req.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "заголовок обязателен")
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	id := a.nextID
	a.nextID++
	a.sightings[id] = &ufo{ID: id, Title: req.GetTitle(), Content: req.GetContent()}

	slog.Info("наблюдение создано", "id", id)

	return &ufo_v1.CreateUfoResponse{Id: id}, nil
}

// GetUfo возвращает наблюдение по ID (gRPC)
func (a *api) GetUfo(_ context.Context, req *ufo_v1.GetUfoRequest) (*ufo_v1.GetUfoResponse, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	n, exists := a.sightings[req.GetId()]
	if !exists {
		return nil, status.Error(codes.NotFound, "наблюдение не найдено")
	}

	return &ufo_v1.GetUfoResponse{Id: n.ID, Title: n.Title, Content: n.Content}, nil
}
