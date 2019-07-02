package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	type IndexInfo struct {
		IsSet bool
		Index int
	}
	var indexMap map[int]IndexInfo
	indexMap = make(map[int]IndexInfo)
	for index, value := range nums {
		otherNumber := target - value
		if indexMap[otherNumber].IsSet {
			return []int{indexMap[otherNumber].Index, index}
		}
		// 一定要放在后面执行,因为会有 [3,3] 6 的情况
		indexMap[value] = IndexInfo {
			IsSet: true,
			Index: index,
		}
	}
	return []int{}
}
func main () {
	// fmt.Print(twoSum([]int{2,7,11,15}, 9))
	// fmt.Print(twoSum([]int{3,2,4}, 6))
	// fmt.Print(twoSum([]int{3,3}, 6))
	//fmt.Print(twoSum([]int{3,2,3}, 6)) // 0 2
	//fmt.Print(twoSum([]int{0,4,3,0}, 0))
	//fmt.Print(twoSum([]int{-1,-2,-3,-4,-5}, -8)) // -2 -4
	fmt.Print(twoSum([]int{-10,7,19,15}, 9)) // 0 2
}
