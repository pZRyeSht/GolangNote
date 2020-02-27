func findWords(words []string) []string {
    res := make([]string,0)
	map1 := map[string]int{"q":1,"w":1,"e":1,"r":1,"t":1,"y":1,"u":1,"i":1,"o":1,"p":1}
	map2 := map[string]int{"a":1,"s":1,"d":1,"f":1,"g":1,"h":1,"j":1,"k":1,"l":1}
	map3 := map[string]int{"z":1,"x":1,"c":1,"v":1,"b":1,"n":1,"m":1,}
	maps := map[int]map[string]int{1: map1, 2: map2, 3: map3}
	for _, v := range words {
		temp := 0
		flag := true
		for k, v1 := range maps {
			if v1[string(v[0])] != 0 || v1[string(v[0] + 32)] != 0 {
				temp = k
				break
			}
		}
		for _, v2 := range v {
			if maps[temp][string(v2)] == 0 && maps[temp][string(v2 + 32)] == 0 {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, v)
		}
	}
	return res
}