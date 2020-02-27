1.
func sortArrayByParity(A []int) []int {
    j := []int{}
    o := []int{}
    for _, v := range A {
        if v % 2 == 0 {
            o = append(o, v)
        } else {
            j = append(j, v)
        }
    }
    res := make([]int, len(o) + len(j))
    for i, v := range o {
        res[i] = v
    }
     for i,v := range j {
        res[i + len(o)] = v
    }
    return res
}
2.
func sortArrayByParity(A []int) []int {
   temp := 0
   for i, v := range A {
       if v % 2 == 0 {
           A[i], A[temp] = A[temp], A[i]
           temp++
       }
   }
    return A
}
3.
func sortArrayByParity(A []int) []int {
    l := len(A)
    res := make([]int, l)
   i := 0
   j := 1
   for _, v := range A {
       if v % 2 == 0 {
           res[i] = v
           i++
       } else {
           res[l - j] = v
           j++
       }
   }
    return res
}