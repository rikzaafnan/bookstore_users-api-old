package date_utils

import (
	"fmt"
	"time"
)

const (
	apiDateLayout = "02-01-2006T15:04:05Z"
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
