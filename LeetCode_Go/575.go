func distributeCandies(candies []int) int {
    candiesMap := make(map[int]int)
    for _, v := range candies {
        candiesMap[v] = 1
    }
    if len(candiesMap) <= len(candies)/2 {
        return len(candiesMap)
    }
    return len(candies)/2
}