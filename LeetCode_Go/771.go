1.
func numJewelsInStones(J string, S string) int {
    count := 0
    jmap := make(map[rune]int)
    for _, v := range J {
        jmap[v] = 1
    }
    for _, v := range S {
        if _, ok := jmap[v]; ok {
            count++
        }
    }
    return count
}
2.
func numJewelsInStones(J string, S string) int {
    count := 0
    for _, v := range S {
        if strings.Contains(J, string(v)) {
             count++
        }
    }
    return count
}