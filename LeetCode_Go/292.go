1.
func canWinNim(n int) bool {
    return n % 4 != 0
}
2.
func canWinNim(n int) bool {
    return n & 3 != 0
}