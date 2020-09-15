package util

import (
	"math/rand"
	"time"
)

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// GenerateRandomString 生成随机字符串
func GenerateRandomString(n int, allowedChars []rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// GenerateRandomArray 生成随机数列
func GenerateRandomArray(n, min, max int) []int {
	array := make([]int, n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		array[i] = randInt(min, max)
	}
	return array
}

// GenerateNearlyOrderedArray 生成基本有序的数列
func GenerateNearlyOrderedArray(n, swapTimes int) []int {
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = i
	}

	for j := 0; j < swapTimes; j++ {
		posx := randInt(0, n-1)
		posy := randInt(0, n-1)
		array[posx], array[posy] = array[posy], array[posx]
	}
	return array
}

// IsSorted 判断是否正序
func IsSorted(list []int) bool {
	len := len(list)
	for i := 0; i < len; i++ {
		if list[i] > list[i+1] {
			return false
		}
	}
	return true
}

// randInt 生成[min, max)范围内随机数
func randInt(min, max int) int {
	if min < 0 || max < 0 {
		return 0
	}
	if min > max {
		return max
	}
	return rand.Intn(max-min) + min
}
