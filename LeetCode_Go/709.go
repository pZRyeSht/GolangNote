func toLowerCase(str string) string {
   res:= ""
	for _, v := range str {
		if v <= 90 && v >= 65 {
			res += string(v+32)
		} else {
			res += string(v)
		}
	}
	return res
}