package valueobject

import (
	"fmt"

	errs "github.com/mbakhodurov/examples2/week_4/ddd/app/internal/errors"
)

// GPUProperties — экспортируемый алиас для приватной структуры
type GPUProperties = gpuProperties

// type GPUProperties interface {
// 	RequiredTDPWatts() int
// 	VRAMGB() int
// }

// gpuProperties — свойства видеокарты (Value Object)
type gpuProperties struct {
	requiredTDPWatts int
	vramGB           int
}

// NewGPUProperties создаёт свойства видеокарты
// RequiredTDP и VRAM должны быть положительными
func NewGPUProperties(requiredTDPWatts, vramGB int) (*ComponentProperties, error) {
	if requiredTDPWatts <= 0 {
		return nil, fmt.Errorf("требуемый TDP должен быть положительным: %w",
			errs.ErrInvalidProperties)
	}
	if vramGB <= 0 {
		return nil, fmt.Errorf("объём VRAM должен быть положительным: %w",
			errs.ErrInvalidProperties)
	}

	return &componentProperties{
		gpu: &gpuProperties{
			requiredTDPWatts: requiredTDPWatts,
			vramGB:           vramGB,
		},
	}, nil
}

// RequiredTDPWatts возвращает минимальный TDP блока питания для видеокарты
func (g *gpuProperties) RequiredTDPWatts() int { return g.requiredTDPWatts }

// VRAMGB возвращает объём видеопамяти в гигабайтах
func (g *gpuProperties) VRAMGB() int { return g.vramGB }
