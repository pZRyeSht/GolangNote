1.
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[j] == target - nums[i]) {
				return []int { i, j }
			}
		}
	}
	return nil
}
2.
func twoSum(nums []int, target int) []int {
    hash := make(map[int]int)
    res := make([]int, 2)
    for k, v := range nums {
        hash[v] = k
    }
    for k ,v := range nums {
        value := target - v
        index, exist := hash[value]
        if exist && index != k {
            res = []int{k, index}
            break
        }
    }
    return res
}