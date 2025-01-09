init:
	cd api/third_party && git clone https://github.com/protocolbuffers/protobuf.git src/google/protobuf
create:
	go run main.go start

split:
	go run main.go ffmpeg split

proto:
	protoc --proto_path=./api \
 	       --go_out=paths=source_relative:./api \
 	       --go-grpc_opt=require_unimplemented_servers=false \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
		   ./api/service/split_video.proto