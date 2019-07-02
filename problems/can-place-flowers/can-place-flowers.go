// https://github.com/nimojs/learn/leetcode/problems/can-place-flowers/
package main

func canPlaceFlowers(flowerbed []int, n int) bool {
  flowerbed = append([]int{0}, flowerbed...)
  flowerbed = append(flowerbed,0)
  var currentRangeEmptyCount int
  canPlaceCount := 0
  flowerbedLen := len(flowerbed)
  for index, item := range flowerbed {
	  if item == 0 {
	  	currentRangeEmptyCount++
	  }
	  if index == flowerbedLen-1 || item == 1 {
		  if currentRangeEmptyCount>=3 {
		  	// floor( (currentRangeEmptyCount-1) / 2 )
		  	// go 语言特性 int / 自动 floor
		  	canPlaceCount = canPlaceCount + (currentRangeEmptyCount - 1) / 2
		  }
	  }
	  if item == 1 {
	  	currentRangeEmptyCount = 0
	  }
  }
  return canPlaceCount >= n
}

func main() {
	print(canPlaceFlowers([]int{1,0,0,0,0,0,0,0,0,1}, 3))
}