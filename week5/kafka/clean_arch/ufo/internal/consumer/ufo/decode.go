package ufo

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"

	events_v1 "github.com/mbakhodurov/examples2/week_5/kafka/clean_arch/shared/pkg/proto/events/v1"
	"github.com/mbakhodurov/examples2/week_5/kafka/clean_arch/ufo/internal/model"
)

func decodeUFORecorded(data []byte) (model.UFORecordedEvent, error) {
	var pb events_v1.UFORecorded
	if err := proto.Unmarshal(data, &pb); err != nil {
		return model.UFORecordedEvent{}, fmt.Errorf("не удалось десериализовать protobuf: %w", err)
	}

	var observedAt *time.Time
	if pb.ObservedAt != nil {
		observedAt = new(pb.ObservedAt.AsTime())
	}

	return model.UFORecordedEvent{
		UUID:        pb.Uuid,
		ObservedAt:  observedAt,
		Location:    pb.Location,
		Description: pb.Description,
	}, nil
}
