init:
	cd api/third_party && git clone https://github.com/protocolbuffers/protobuf.git src/google/protobuf
create_local:
	go run main.go storage create_local

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

combine:
	go run main.go ffmpeg combine-videos -f1=output/video_07_01_25/1197802-hd_1920_1080_25fps.mp4 -f2=output/video_07_01_25/1197802-hd_1920_1080_25fps.mp4 -o=./output/video_07_01_25/combine

upload:
	go run main.go youtube upload -f=./output/watermark/test.mp4 -t=test -d=test

oauth: 
	go run main.go youtube oauth2

gpt:
	go run main.go gpt summarize -subject="rick and morty" -episode="season 1 episode 7"

download:
	go run main.go youtube download -url="https://www.youtube.com/watch?v=3Bq8uokFqOo"

extract_audio:
	go run main.go ffmpeg extract_audio

init_source:
	go run main.go youtube list

full:
	go run main.go full 
	# -source="https://www.youtube.com/watch?v=3Bq8uokFqOo"

tidy:
	go mod tidy && go mod vendor
