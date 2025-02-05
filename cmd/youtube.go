package cmd

import (
	"github.com/r4g3ch33m5/ffmpeg_video/cmd/youtube"
	"github.com/urfave/cli/v3"
)

// YoutubeCommand defines the main YouTube command group
var YoutubeCommand = &cli.Command{
	Name:  "youtube",
	Usage: "Perform various operations on YouTube videos",
	Commands: []*cli.Command{
		youtube.DownloadCommand,
		youtube.ExtractAudioCommand,
		youtube.MetadataCommand,
		youtube.UploadCommand,
		youtube.Oauth2Command,
		youtube.ListVideoCommand,
	},
}
