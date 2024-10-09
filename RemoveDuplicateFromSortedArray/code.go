package removeduplicatefromsortedarray

func removeDuplicates(nums []int) int {
	for i, num := range nums {
		if i + 1 < len(nums) && num == nums[i + 1] {
			nums := append(nums[:i], nums[i+1:]...)
			return removeDuplicates(nums)
		}
	}
	return len(nums)
}
