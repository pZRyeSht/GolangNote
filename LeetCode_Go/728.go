1.
func selfDividingNumbers(left int, right int) []int {
    number := make([]int, 0)
    for i := left; i <= right; i++ {
        temp := func(int) bool{
            x:= i
            for x > 0 {
                y := x % 10
                x /= 10
                if y == 0 || i % y > 0 {
                    return false
                }
            }
            return true
        }(i)
        if temp {
            number = append(number, i)
        }
    }
    return number
}
2.
func selfDividingNumbers(left int, right int) []int {
	number := make([]int, 0)
	for i := left; i <= right; i++ {
		temp := func(int) bool{
			numstr := strconv.Itoa(i)
			for _, v := range numstr {
				if v == '0' || int32(i) % (v - '0') > 0 {
					return false
				}
			}
			return true
		}(i)
		if temp {
			number = append(number, i)
		}
	}
	return number
}
3.
func selfDividingNumbers(left int, right int) []int {
	number := make([]int, 0)
	for i := left; i <= right; i++ {
		if selfDividing(i) {
			number = append(number, i)
		}
	}
	return number
}
func selfDividing(num int) bool {
	x := num
	for x > 0 {
		y := x % 10
		x /= 10
		if y == 0 || num % y > 0 {
			return false
		}
	}
	return true
}