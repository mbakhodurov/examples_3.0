package ufo

import (
	"sync"

	"github.com/mbakhodurov/examples2/week_2/layers/internal/repository/record"
)

type repository struct {
	mu   sync.RWMutex
	data map[string]record.Sighting
}

func NewRepository() *repository {
	return &repository{
		data: make(map[string]record.Sighting),
	}
}
