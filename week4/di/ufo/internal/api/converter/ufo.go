// Package converter содержит функции преобразования между транспортными DTO
// и доменными моделями сервиса наблюдений НЛО
package converter

import (
	"time"

	ufo_v1 "github.com/mbakhodurov/examples2/week_4/di/shared/pkg/proto/ufo/v1"
	"github.com/mbakhodurov/examples2/week_4/di/ufo/internal/model"
	"github.com/mbakhodurov/examples2/week_4/di/ufo/internal/service/input"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// CreateRequestToInput преобразует запрос создания наблюдения во вход use case'а
func CreateRequestToInput(req *ufo_v1.CreateRequest) input.CreateSightingInput {
	var observedAt *time.Time
	if req.ObservedAt != nil {
		observedAt = lo.ToPtr(req.ObservedAt.AsTime())
	}

	var color *string
	if req.Color != nil {
		color = lo.ToPtr(req.Color.Value)
	}

	var sound *bool
	if req.Sound != nil {
		sound = lo.ToPtr(req.Sound.Value)
	}

	var durationSeconds *int32
	if req.DurationSeconds != nil {
		durationSeconds = lo.ToPtr(req.DurationSeconds.Value)
	}

	return input.CreateSightingInput{
		ObservedAt:      observedAt,
		Location:        req.Location,
		Description:     req.Description,
		Color:           color,
		Sound:           sound,
		DurationSeconds: durationSeconds,
	}
}

// UpdateRequestToInput преобразует запрос обновления наблюдения во вход use case'а
func UpdateRequestToInput(req *ufo_v1.UpdateRequest) input.UpdateSightingInput {
	var observedAt *time.Time
	if req.ObservedAt != nil {
		observedAt = lo.ToPtr(req.ObservedAt.AsTime())
	}

	var location *string
	if req.Location != nil {
		location = lo.ToPtr(req.Location.Value)
	}

	var description *string
	if req.Description != nil {
		description = lo.ToPtr(req.Description.Value)
	}

	var color *string
	if req.Color != nil {
		color = lo.ToPtr(req.Color.Value)
	}

	var sound *bool
	if req.Sound != nil {
		sound = lo.ToPtr(req.Sound.Value)
	}

	var durationSeconds *int32
	if req.DurationSeconds != nil {
		durationSeconds = lo.ToPtr(req.DurationSeconds.Value)
	}

	return input.UpdateSightingInput{
		ObservedAt:      observedAt,
		Location:        location,
		Description:     description,
		Color:           color,
		Sound:           sound,
		DurationSeconds: durationSeconds,
	}
}

// SightingToDTO конвертирует доменную модель наблюдения в транспортный DTO
func SightingToDTO(s model.Sighting) *ufo_v1.Sighting {
	var observedAt *timestamppb.Timestamp
	if s.ObservedAt != nil {
		observedAt = timestamppb.New(*s.ObservedAt)
	}

	var color *wrapperspb.StringValue
	if s.Color != nil {
		color = wrapperspb.String(*s.Color)
	}

	var sound *wrapperspb.BoolValue
	if s.Sound != nil {
		sound = wrapperspb.Bool(*s.Sound)
	}

	var durationSeconds *wrapperspb.Int32Value
	if s.DurationSeconds != nil {
		durationSeconds = wrapperspb.Int32(*s.DurationSeconds)
	}

	var updatedAt *timestamppb.Timestamp
	if s.UpdatedAt != nil {
		updatedAt = timestamppb.New(*s.UpdatedAt)
	}

	var deletedAt *timestamppb.Timestamp
	if s.DeletedAt != nil {
		deletedAt = timestamppb.New(*s.DeletedAt)
	}

	return &ufo_v1.Sighting{
		Uuid:            s.Uuid,
		ObservedAt:      observedAt,
		Location:        s.Location,
		Description:     s.Description,
		Color:           color,
		Sound:           sound,
		DurationSeconds: durationSeconds,
		CreatedAt:       timestamppb.New(s.CreatedAt),
		UpdatedAt:       updatedAt,
		DeletedAt:       deletedAt,
	}
}
