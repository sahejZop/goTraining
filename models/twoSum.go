package models

func TwoSum(nums []int, target int) []int {
	var result []int
	mapCount := map[int]int{}
	for i := 0; i < len(nums); i++ {
		value, itExists := mapCount[target-nums[i]]
		if itExists {
			//result = append(result, value, i)
			//return result
			return []int{value, i}
		} else {
			mapCount[nums[i]] = i
		}
	}

	return result
}
