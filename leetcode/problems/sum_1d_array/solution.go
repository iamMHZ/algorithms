package sum_1d_array

func RunningSum(nums []int) []int {
	length := len(nums)
	resultSlice := make([]int, length)

	for i, value := range nums {
		if i == 0 {
			resultSlice[i] = value
		} else {
			resultSlice[i] = value + resultSlice[i-1]
		}

	}

	return resultSlice
}

// it can be solved with one array.
func RunningSumInPlace(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}

	return nums
}
