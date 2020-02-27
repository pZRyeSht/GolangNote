1.
func defangIPaddr(address string) string {
   return(strings.Replace(address,".","[.]",-1))
}
2.
func defangIPaddr(address string) string {
    res := ""
    for _, v := range address {
        temp := "[.]"
        if string(v) == "." {
            res += temp
        } else {
            res += string(v)
        }
    }
    return res
}