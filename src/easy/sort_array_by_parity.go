package easy

// https://leetcode.com/problems/sort-array-by-parity/

func sortArrayByParity(A []int) []int {
	var evenArray []int
	var oddArray []int

	for _, v := range A {
		if v%2 == 0 {
			evenArray = append(evenArray, v)
		} else {
			oddArray = append(oddArray, v)
		}
	}
	evenArray = append(evenArray, oddArray...)
	return evenArray
}
