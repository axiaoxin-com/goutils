package goutils

import (
	"hash/maphash"
	"math/rand"
)

// RandInt 返回一个范围为 [start, end] 的int型随机数
func RandInt(start, end int) int {
	r := rand.New(rand.NewSource(int64(new(maphash.Hash).Sum64())))
	num := start + r.Intn(end-start+1)
	return num
}
