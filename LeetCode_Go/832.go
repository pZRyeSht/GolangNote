1.
func flipAndInvertImage(A [][]int) [][]int {
    for _, row := range A {
        for i := len(row)/2-1; i >= 0; i-- {
            opp := len(row)-1-i
            row[i], row[opp] = row[opp], row[i]
        }
        for j, v := range row {
            row[j] = 1 - v
        }
    }
    return A
}
2.
func flipAndInvertImage(A [][]int) [][]int {
    for _, a := range A {
       for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
           a[left], a[right] = a[right], a[left]
        }
        for j, v := range a {
            a[j] = 1 - v
        }
    }
    return A
}