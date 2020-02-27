func game(guess []int, answer []int) int {
    count := 0
    for k, v := range guess {
        if v == answer[k] {
            count++
        }
    }
    return count
}