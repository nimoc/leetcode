package main

import (
	"math"
)
func reverse(x int) (o int) {
	for x != 0 {
		o = o * 10 // 进一位
		lastInt :=  x % 10 // 获取 x 最后一个 int
		o += lastInt // o push
		x = x / 10  // x shift
	}
	if o < math.MinInt32 || o > math.MaxInt32 {
		return 0
	}
	return o
}
func main () {
	reverse(123)
}