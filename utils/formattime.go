package utils

import (
	"strings"
	"time"
)

var (
	Now             = time.Now()
	DateOnly        = "2006-01-02"
	TimeOnly        = "15:04:05"
	DateTime        = "2006-01-02 15:04:05"
	Today           = time.Now().Format(DateOnly)
	Yesterday       = time.Now().AddDate(0, 0, -1).Format(DateOnly)
	CurrentDateTime = Now.Format(DateTime)
	StartOfDateTime = time.Date(Now.Year(), Now.Month(), Now.Day(), 0, 0, 0, 0, Now.Location()).Format(DateTime)
)

func ExtractTime(diffInText string) (timeInText string) {
	result := strings.Split(diffInText, ":")
	return result[0] + " jam " + result[1] + "menit "
}
