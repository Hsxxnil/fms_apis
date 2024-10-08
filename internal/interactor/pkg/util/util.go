package util

import (
	"math"
	"math/rand"
	"time"
)

func PointerString(s string) *string     { return &s }
func PointerInt64(i int64) *int64        { return &i }
func PointerBool(b bool) *bool           { return &b }
func PointerTime(t time.Time) *time.Time { return &t }

func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}

func Round(x float64) int64 {
	return int64(math.Floor(x + 0.5))
}

/*
Used: 移除字串陣列指定的項目
Parameters: []string, string
*/
func RemoveString(s []string, item string) []string {
	for i, v := range s {
		if v == item {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
