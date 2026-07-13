package config

type config struct {
	GRPC grpcConfig `yaml:"grpc"`
	PG   pgConfig   `yaml:"pg"`
}

const defaultConfigPath = "config.local.yaml"
