module github.com/mbakhodurov/examples2/week_6/easy_auth/http_frontend

go 1.25.4

replace github.com/mbakhodurov/examples2/week_6/easy_auth/shared => ../shared

require (
	github.com/go-chi/chi/v5 v5.3.1
	github.com/mbakhodurov/examples2/week_6/easy_auth/shared v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.83.0
)

require (
	golang.org/x/net v0.55.0 // indirect
	golang.org/x/sys v0.45.0 // indirect
	golang.org/x/text v0.37.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260526163538-3dc84a4a5aaa // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
