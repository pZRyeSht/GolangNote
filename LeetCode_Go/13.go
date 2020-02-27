func romanToInt(s string) int {
    map1 := map[string]int{
		"I":1,
		"V":5,
		"X":10,
		"L":50,
		"C":100,
		"D":500,
		"M":1000,
	}
	map2 := map[string]int{
		"IV":4,
        "IX":9,
        "XL":40,
        "XC":90,
        "CD":400,
        "CM":900,
	}

	var rev []int
	var number int
	for index,val := range map2{
		if strings.Contains(s,index){
			rev = append(rev,val)
			s = strings.Replace(s,index,"",1)
		}
	}

	for _,val := range s{
		if map1[string(val)] != 0{
			rev = append(rev,map1[string(val)])
		}
	}
		

	for _,value := range rev{
		number += value
	}
	return number
}