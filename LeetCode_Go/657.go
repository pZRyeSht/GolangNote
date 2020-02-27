1.
func judgeCircle(moves string) bool {
    up := 0
    left := 0
    for _, v := range moves {
        if string(v) == "U" {
            up ++
        } else if string(v) == "D" {
            up--
        } else if string(v) == "L" {
            left++
        } else {
            left--
        }
    }
    return up == 0 && left == 0
}
2.
func judgeCircle(moves string) bool {
    up := 0
    left := 0
    for _, v := range moves {
        if v == 85 {
            up ++
        } else if v == 68 {
            up--
        } else if v == 76 {
            left++
        } else {
            left--
        }
    }
    return up == 0 && left == 0
}