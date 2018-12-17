package easy

// https://leetcode.com/problems/unique-email-addresses/
// not finished
// need to leanrn string, byte,rune difference

import "fmt"

func numUniqueEmails(emails []string) int {
	var ret = make(map[string]int)
	for i := 0; i < len(emails); i++ {

		s2 := []rune{}
		var flag = true
		for _, c := range emails[i] {
			if c == '.' {
				flag = true
			} else if c == '+' {
				flag = false
			} else {
				if flag {
					s2 = append(s2, c)
				}
			}
		}

		fmt.Println(string(s2))
		ret[string(s2)] = 1
	}

	return len(ret)
}
