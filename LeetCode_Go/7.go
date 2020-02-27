func reverse(x int) int {
    var rev int
    for x != 0 {
        rev = rev * 10 + x % 10
        x = x / 10
    }
    if rev > math.MaxInt32 || rev < math.MinInt32 {
        return 0
    }
    return rev
}