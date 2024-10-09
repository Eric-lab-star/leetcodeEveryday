package twosum;

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int);
	for i, num := range nums {
		comp := target - num;
		if idx, ok := numMap[comp]; ok {
			return []int{idx, i};
		}
		numMap[num] = i
	}
	return []int{};
}

