package valueobject

import errs "github.com/mbakhodurov/examples2/week_4/ddd/app/internal/errors"

// ComponentType — тип комплектующего (перечисление)
type ComponentType string

const (
	// ComponentTypeMotherboard — материнская плата
	ComponentTypeMotherboard ComponentType = "motherboard"
	// ComponentTypeCPU — процессор
	ComponentTypeCPU ComponentType = "cpu"
	// ComponentTypeRAM — оперативная память
	ComponentTypeRAM ComponentType = "ram"
	// ComponentTypeGPU — видеокарта
	ComponentTypeGPU ComponentType = "gpu"
)

// NewComponentType создаёт ComponentType с валидацией допустимых значений
func NewComponentType(s string) (ComponentType, error) {
	ct := ComponentType(s)

	switch ct {
	case ComponentTypeMotherboard, ComponentTypeCPU, ComponentTypeRAM, ComponentTypeGPU:
		return ct, nil
	default:
		return "", errs.ErrInvalidComponentType
	}
}
