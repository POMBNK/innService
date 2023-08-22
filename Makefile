.PHONY: all

proto:gen_model gen_service

gen_model:
	protoc --proto_path=proto --go_out=gen/go/ --go_opt=paths=source_relative \
		   --go-grpc_out=gen/go --go-grpc_opt=paths=source_relative \
		   proto/inn_service/model/v1/*.proto
gen_service:
	protoc --proto_path=proto --go_out=gen/go/ --go_opt=paths=source_relative \
		   --go-grpc_out=gen/go --go-grpc_opt=paths=source_relative \
		   proto/inn_service/service/v1/*.proto

