package tools

import (
	"fmt"
	"go-devops-mimi/server/config"
	"time"
)

func String2Time(timeStr string) time.Time {
	loc, _ := time.LoadLocation(config.Conf.System.TimeZone)

	if len(timeStr) == 10 {
		timeStr = fmt.Sprintf("%s 00:00:00", timeStr)
	}

	checkTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	return checkTime
}
func GetTimeFormat(t *time.Time) string {
	if t == nil {
		return ""
	} else {
		return t.Format(time.DateTime)
	}
}
