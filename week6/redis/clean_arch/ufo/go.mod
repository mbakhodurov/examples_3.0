module github.com/mbakhodurov/examples2/week_6/redis/clean_arch/ufo

replace github.com/mbakhodurov/examples2/week_6/redis/clean_arch/shared => ../shared

replace github.com/mbakhodurov/examples2/week_6/redis/clean_arch/platform => ../platform

go 1.25.4

require (
	github.com/google/uuid v1.6.0
	github.com/ilyakaznacheev/cleanenv v1.5.0
	github.com/jackc/pgx/v5 v5.10.0
	github.com/joho/godotenv v1.5.1
	github.com/mbakhodurov/examples2/week_6/redis/clean_arch/platform v0.0.0-00010101000000-000000000000
	github.com/mbakhodurov/examples2/week_6/redis/clean_arch/shared v0.0.0-00010101000000-000000000000
	github.com/redis/go-redis/v9 v9.22.0
	github.com/samber/lo v1.53.0
	google.golang.org/grpc v1.83.0
	google.golang.org/protobuf v1.36.11
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.16.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/net v0.55.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/sys v0.45.0 // indirect
	golang.org/x/text v0.37.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260526163538-3dc84a4a5aaa // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
