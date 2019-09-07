// https://github.com/nimojs/learn/leetcode/problems/can-place-flowers/
function canPlaceFlowers(flowerbed, n) {
  if (n===0) return true
	flowerbed.push(0)
	flowerbed.unshift(0)
	let currentRangeEmptyCount = 0
	let canPlaceCount = 0
	let flowerbedLen = flowerbed.length
	flowerbed.forEach(function (item, index) {
	   if (item === 0) {
	  	currentRangeEmptyCount++
	  }
	  if (index === flowerbedLen-1 || item === 1) {
		  if (currentRangeEmptyCount >= 3) {
		  	canPlaceCount = canPlaceCount + Math.floor((currentRangeEmptyCount-1)/2)
		  }
	  }
	  if (item === 1){
	  	currentRangeEmptyCount = 0
	  }
	})
	return canPlaceCount >= n
}