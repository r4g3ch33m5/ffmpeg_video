package util

import "time"

func GetTodayFolder() string {
	return "video_" + time.Now().Format("02_01_06")
}
