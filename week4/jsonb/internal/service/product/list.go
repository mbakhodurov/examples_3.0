package product

import (
	"context"
	"fmt"

	"github.com/mbakhodurov/examples2/week_4/jsonb/internal/model"
)

// List возвращает все товары
func (s *Service) List(ctx context.Context) ([]model.Product, error) {
	products, err := s.productRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("получить список товаров: %w", err)
	}

	return products, nil
}
