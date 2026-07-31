package ufo

import (
	"context"

	"github.com/mbakhodurov/examples2/week_5/kafka/clean_arch/platform/pkg/kafka"
)

type KafkaProducer interface {
	Send(ctx context.Context, msg *kafka.Message) error
}
