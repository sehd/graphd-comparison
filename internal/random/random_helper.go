package random

import (
	"math/rand"
	"strings"
	"time"
)

func RandomWord(lengthFrom int, lengthTo int) string {
	var sb strings.Builder
	rand.Seed(time.Now().UnixNano())
	length := lengthFrom + rand.Intn(lengthTo-lengthFrom+1)
	for i := 0; i < length; i++ {
		sb.WriteRune(rune('a' + rand.Intn(26)))
	}
	return sb.String()
}

func GetRandomSampleWithReplacement(len int, count int) []int {
	res := make([]int, count)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		res = append(res, rand.Intn(len))
	}
	return res
}
