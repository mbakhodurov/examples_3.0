module github.com/mbakhodurov/examples2/week_1/workspace/grpc

go 1.25.0

replace github.com/mbakhodurov/examples2/week_1/workspace/shared => ../shared

require (
	github.com/google/uuid v1.6.0
	github.com/mbakhodurov/examples2/week_1/workspace/shared v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.81.1
	google.golang.org/protobuf v1.36.11
)

require (
	golang.org/x/net v0.56.0 // indirect
	golang.org/x/sys v0.46.0 // indirect
	golang.org/x/text v0.38.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260226221140-a57be14db171 // indirect
)
