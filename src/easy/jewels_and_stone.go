package main

// https://leetcode.com/problems/jewels-and-stones/

func numJewelsInStones(J string, S string) int {
	jDict := make(map[rune]int)
	ret := 0
	for _, v := range J {
		jDict[v] = 1
	}

	for _, s := range S {
		if _,ok := jDict[s]; ok {
			ret++
		}
	}

	return ret
}