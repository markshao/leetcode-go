package easy

func judgeCircle(moves string) bool {
	var direMap = make(map[rune]int)
	for _, v := range moves {
		_, ok := direMap[v]
		if !ok {
			direMap[v] = 0
		} else {
			direMap[v] = direMap[v] + 1
		}
	}
	if direMap['U']-direMap['D'] != 0 {
		return false
	} else if direMap['L']-direMap['R'] != 0 {
		return false
	}

	return true
}
