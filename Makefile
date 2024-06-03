.PHONY: protoc/auth
protoc/auth:
	@echo 'Updating proto...'
	protoc -I protos protos/auth.proto --go_out=services/auth/gen/auth --go_opt=paths=source_relative --go-grpc_out=services/auth/gen/auth --go-grpc_opt=paths=source_relative