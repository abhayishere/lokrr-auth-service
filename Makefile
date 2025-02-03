protoc:
	cd pb && protoc -I . \
    --grpc-gateway_out ../proto/ \
    --grpc-gateway_opt paths=source_relative \
    **/*.proto && \
    protoc -I . \
    --go_out ../proto/ --go_opt paths=source_relative \
    --go-grpc_out ../proto/ --go-grpc_opt paths=source_relative \
    **/*.proto

docker-image:
    docker build -t auth-service:latest

docker-run:
    docker run -p 50052:50052 auth-service:latest