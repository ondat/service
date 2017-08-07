# dataplane service protobuf 

go:
	protoc -I dataplane/ dataplane/dataplane.proto --go_out=plugins=grpc:dataplane