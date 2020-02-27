1.
func candy(ratings []int) int {
    if ratings == nil || len(ratings) == 0 {
        return 0
    }
    lens := len(ratings)
    leftToRight := make([]int,lens)
    rightToLeft := make([]int,lens)
    sum := 0
    for i := 0; i < lens; i++ {
        leftToRight[i] = 1
        rightToLeft[i] = 1
    }
    for i := 1; i < lens; i++ {
        if ratings[i] > ratings[i-1] {
            leftToRight[i] = leftToRight[i-1] + 1
        }
    }
    for i := lens-2; i >= 0; i-- {
        if ratings[i] > ratings[i+1] {
            rightToLeft[i] = rightToLeft[i+1] + 1
        }
    }
    for i := 0; i < lens; i++ {   
        if leftToRight[i] >= rightToLeft[i] {
            sum += leftToRight[i]
        } else {
            sum += rightToLeft[i]
        }
    }
    return sum
}
2.
func candy(ratings []int) int {
    if ratings == nil || len(ratings) == 0 {
		return 0
	}
	lens := len(ratings)
	candies := make([]int, lens)
	sum := 0
	flag := true
	for i := 0; i < lens; i++ {
		candies[i] = 1
	}
	for flag {
		flag = false
		for i := 0; i < lens; i++ {
			if i != lens -1 && ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
				candies[i] = candies[i+1] + 1
				flag = true
			}
			if i > 0 && ratings[i] > ratings[i-1] && candies[i] <= candies[i-1] {
				candies[i] = candies[i-1] + 1
				flag = true
			}
		}
	}
	for k := range candies {
		sum += candies[k]
	}
	return sum
}
3.
func candy(ratings []int) int {
    if ratings == nil || len(ratings) == 0 {
		return 0
	}
	lens := len(ratings)
	candies := make([]float64, lens)
	for k := range candies {
		candies[k] = 1
	}
	for i := 1; i < lens; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	sum := candies[lens-1]
	for i := lens-2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = math.Max(candies[i], candies[i+1] + 1)
		}
		sum += candies[i]
	}
	return int(sum)
}