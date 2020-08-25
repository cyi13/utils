package rands

import (
	"math/rand"
	"time"

	"github.com/guiychen/snippet/funcs"
)

//RandInts 生成指定区间的一系列随机数
func RandInts(start, end int, count int) []int {
	nums := []int{}
	if end < start {
		return nums
	}
	ints := []int{}
	for i := start; i < end; i++ {
		ints = append(ints, i)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	maxRange := len(ints)
	for i := 0; i < maxRange; i++ {
		mLen := len(ints)
		n := r.Intn(mLen)
		nums = append(nums, ints[n])
		ints = funcs.DeleteInts(ints, n)
		if len(nums) >= count {
			break
		}
	}
	return nums
}

var str = []rune("abcdefghijklmnopqrstuvwxyz")

func RandString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = str[rand.Intn(len(str))]
	}
	return string(b)
}
