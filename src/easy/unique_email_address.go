package main

// https://leetcode.com/problems/unique-email-addresses/
// not finished
// need to leanrn string, byte,rune difference

func numUniqueEmails(emails []string) int {
	var ret map[rune]int = make(map[rune]int)
	for i:=0;i<len(emails);i++ {
		
		var s2 rune
		var flag = true
		for _,c := range emails[i] {
			if c == '.' {
				flag = true
			} else if c == '+' {
				flag = false
			} else {
				if flag {
					s2 = s2 + c
				}
			}
		}
		ret[s2] = 1
	}

	return len(ret) 
}