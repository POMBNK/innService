proto:
	rm -rf gen/go/*
# create model
	protoc --proto_path=proto --go_out=gen/go/ --go_opt=paths=source_relative \
		   --go-grpc_out=gen/go --go-grpc_opt=paths=source_relative \
		   proto/inn_service/model/v1/*.proto
#create service
	protoc --proto_path=proto --go_out=gen/go/ --go_opt=paths=source_relative \
		   --go-grpc_out=gen/go --go-grpc_opt=paths=source_relative \
		   --grpc-gateway_out=gen/go/ --grpc-gateway_opt=paths=source_relative \
		   proto/inn_service/service/v1/*.proto

.PHONY: proto 