init:
	cd api/third_party && git clone https://github.com/protocolbuffers/protobuf.git src/google/protobuf
create:
	go run main.go start

split:
	go run main.go ffmpeg split

watermark: 
	go run main.go ffmpeg watermark -input=./source/video_07_01_25/1197802-hd_1920_1080_25fps.mp4 -output=./output/watermark/test.mp4

proto:
	protoc --proto_path=./api \
 	       --go_out=paths=source_relative:./api \
 	       --go-grpc_opt=require_unimplemented_servers=false \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
		   ./api/service/split_video.proto

tidy: 
	go mod tidy && go mod vendor 