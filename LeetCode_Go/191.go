1.
func hammingWeight(num uint32) int {
    res := 0
	for num != 0 {
		res ++
		num = num & (num-1)
	}
	return res
}
2.
func hammingWeight(num uint32) int {
    res := 0
	var bits uint32 = 1
	for i := 0; i < 32; i++ {
		if (num&bits) != 0 {
			res ++
		}
		bits = bits << 1
	}
	return res
}