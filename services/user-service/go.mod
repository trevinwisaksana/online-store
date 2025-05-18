module github.com/trevinwisaksana/online-store/user-service

go 1.23.0

require (
	github.com/trevinwisaksana/online-store/proto v0.0.0
	google.golang.org/grpc v1.72.1
)

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.3 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250512202823-5a2f75b736a9 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250505200425-f936aa4a68b2 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

replace github.com/trevinwisaksana/online-store/proto => ../../proto
