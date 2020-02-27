1.
func totalHammingDistance(nums []int) int {
    res := 0
	for k,v := range nums {
		for i:= k+1; i < len(nums); i++ {
			res += bits.OnesCount(uint(v) ^ uint(nums[i]))
		}
	}
	return res
}
2.
func totalHammingDistance(nums []int) int {
    res, n := 0, 0
    for i := uint32(0); i < 32; i++ {
        n = 0
        bits := 1<<i
        for _, v := range nums {
            if v&bits != 0 {
                n++
            }
        }
        res += n * (len(nums) - n)
    }
    return res
}