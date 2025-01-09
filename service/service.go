package service

import (
	"context"

	ffmpeg_video "github.com/r4g3ch33m5/ffmpeg_video/api/service"
)

type Service interface {
	FfmpegService
}

type FfmpegService interface {
	SplitVideo(ctx context.Context, req *ffmpeg_video.SplitVideoRequest) error
}

type service struct {
	*ffmpegServiceImpl
}

func NewService() Service {
	return service{
		ffmpegServiceImpl: &ffmpegServiceImpl{},
	}
}

func (s *ffmpegServiceImpl) SplitVideo(ctx context.Context, req *ffmpeg_video.SplitVideoRequest) error {
	return SplitVideoIntoChunks(ctx, req)
}
