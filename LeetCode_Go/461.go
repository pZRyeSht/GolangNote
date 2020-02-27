1.
func hammingDistance(x int, y int) int {
    return bits.OnesCount(uint(x) ^ uint(y))
}
2.
func hammingDistance(x int, y int) int {
    res := 0
	n := x ^ y
	for n != 0 {
		res++
		n = n & (n-1)
	}
	return res
}
