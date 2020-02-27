func isPalindrome(x int) bool {
    num := x
    var rev int
    for num != 0 {
        rev = rev * 10 + num % 10
        num = num / 10
    }
    if x < 0 || rev != x {
        return false
    }
    return true
}