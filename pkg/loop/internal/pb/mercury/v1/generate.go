// NOTE: the relative paths in the proto_path are to ensure we find common utilities, like BigInt
//go:generate protoc --proto_path=.:../.. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative datasource_v1.proto
//go:generate protoc --proto_path=.:../.. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative reportcodec_v1.proto

package mercury_v1_pb
