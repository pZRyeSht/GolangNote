//贪心算法
func balancedStringSplit(s string) int {
    sum := 0
    res := 0
    for _, v := range s {
        if string(v) == "R" {
            sum ++
        } else {
            sum --
        }
        if sum == 0 {
            res ++
        }
    }
    return res
}