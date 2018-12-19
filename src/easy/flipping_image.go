package easy

//https://leetcode.com/problems/flipping-an-image/

func FlipAndInvertImage(A [][]int) [][]int {
	for _, v := range A {
		length := len(v)
		var i = 0
		for j := length - 1; i < j; i, j = i+1, j-1 {
			v[i], v[j] = v[j], v[i]
			v[i] = 1 ^ v[i]
			v[j] = 1 ^ v[j]
		}
		v[i] = 1 ^ v[i]
	}

	return A
}
