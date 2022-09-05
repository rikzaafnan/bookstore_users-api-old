package date_utils

import (
	"fmt"
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDBLAyout   = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	// ini dengan UTC
	now := time.Now().UTC()

	// ini coba jam utc ke local
	cobKeLocalDariNow := now.Add(time.Hour * 7)
	fmt.Println("=======================")
	fmt.Println(cobKeLocalDariNow)
	fmt.Println("=======================")

	return GetNow().Format(apiDateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDBLAyout)
}
