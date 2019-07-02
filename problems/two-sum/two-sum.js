function twoSum(nums, target) {
	let indexMap = {}
	let output = []
	nums.some(function (value, index) {
		let otherNumber = target - value
		if (indexMap[otherNumber] !== undefined) {
			output = [indexMap[otherNumber], index]
			return true
		}
		indexMap[value] = index
	})
	return output
}