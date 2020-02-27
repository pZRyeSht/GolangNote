func heightChecker(heights []int) int {
    count := 0
    hc := make([]int, 101)
    for _, v := range heights {
		hc[v]++
	}
    j := 0
    for i := 1;i < len(hc); i++ {
		for hc[i] > 0 {
			if heights[j] != i {
				count++
			}
			j++
			hc[i]--
		}
	}
    return count
}