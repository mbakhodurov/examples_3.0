package ufo_cache

import (
	"context"
	"time"

	"github.com/mbakhodurov/examples2/week_6/redis/clean_arch/ufo/internal/model"
	repoConverter "github.com/mbakhodurov/examples2/week_6/redis/clean_arch/ufo/internal/repository/converter"
	"github.com/samber/lo"
)

func (r *repository) Set(ctx context.Context, uuid string, sighting model.Sighting, ttl time.Duration) error {
	cacheKey := r.getCacheKey(uuid)
	err := r.client.HSet(ctx, cacheKey, lo.ToPtr(repoConverter.SightingToRedisView(sighting))).Err()
	if err != nil {
		return err
	}

	return r.client.Expire(ctx, cacheKey, ttl).Err()
}
