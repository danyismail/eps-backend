package utils

import (
	"strings"
	"time"
)

var (
	DateTime  = "2006-01-02 15:04:05"
	DateOnly  = "2006-01-02"
	TimeOnly  = "15:04:05"
	StartDate = time.Now().AddDate(0, 0, -1).Format(DateOnly)
	EndDate   = time.Now().Format(DateOnly)
)

func ExtractTime(diffInText string) (timeInText string) {
	result := strings.Split(diffInText, ":")
	return result[0] + " jam " + result[1] + "menit "
}
