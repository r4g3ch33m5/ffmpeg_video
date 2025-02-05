package cmd

import (
	"path/filepath"

	"github.com/r4g3ch33m5/ffmpeg_video/util"
)

type Record struct{}

var db = map[string]Record{}

func init() {
	util.ApplyAllFileInDir("source", func(path string) error {
		db[filepath.Base(path)] = Record{}
		return nil
	})
}
