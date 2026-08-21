package converter

import (
	"time"

	"github.com/mbakhodurov/examples2/week_6/redis/clean_arch/ufo/internal/model"
	repoModel "github.com/mbakhodurov/examples2/week_6/redis/clean_arch/ufo/internal/repository/model"
	"github.com/samber/lo"
)

// SightingToRedisView — конвертер из модели домена в Redis view
func SightingToRedisView(s model.Sighting) repoModel.SightingRedisView {
	view := repoModel.SightingRedisView{
		UUID:        s.Uuid,
		Location:    s.Location,
		Description: s.Description,
		Color:       s.Color,
		Sound:       s.Sound,
		Duration:    s.DurationSeconds,
		CreatedAtNs: s.CreatedAt.UnixNano(),
	}

	if s.ObservedAt != nil {
		view.ObservedAtNs = lo.ToPtr(s.ObservedAt.UnixNano())
	}
	if s.UpdatedAt != nil {
		view.UpdatedAtNs = lo.ToPtr(s.UpdatedAt.UnixNano())
	}
	if s.DeletedAt != nil {
		view.DeletedAtNs = lo.ToPtr(s.DeletedAt.UnixNano())
	}

	return view
}

// SightingFromPGView - конвертер из PG view в модель домена
func SightingFromPGView(pgView repoModel.SightingPGView) model.Sighting {
	return model.Sighting{
		Uuid:            pgView.UUID,
		ObservedAt:      pgView.ObservedAt,
		Location:        pgView.Location,
		Description:     pgView.Description,
		Color:           pgView.Color,
		Sound:           pgView.Sound,
		DurationSeconds: pgView.DurationSeconds,
		CreatedAt:       pgView.CreatedAt,
		UpdatedAt:       pgView.UpdatedAt,
		DeletedAt:       pgView.DeletedAt,
	}
}

// SightingFromRedisView — конвертер из Redis view в модель домена
func SightingFromRedisView(redisView repoModel.SightingRedisView) model.Sighting {
	sighting := model.Sighting{
		Uuid:            redisView.UUID,
		Location:        redisView.Location,
		Description:     redisView.Description,
		Color:           redisView.Color,
		Sound:           redisView.Sound,
		DurationSeconds: redisView.Duration,
		CreatedAt:       time.Unix(0, redisView.CreatedAtNs),
	}

	if redisView.ObservedAtNs != nil {
		sighting.ObservedAt = lo.ToPtr(time.Unix(0, *redisView.ObservedAtNs))
	}
	if redisView.UpdatedAtNs != nil {
		sighting.UpdatedAt = lo.ToPtr(time.Unix(0, *redisView.UpdatedAtNs))
	}
	if redisView.DeletedAtNs != nil {
		sighting.DeletedAt = lo.ToPtr(time.Unix(0, *redisView.DeletedAtNs))
	}

	return sighting
}
