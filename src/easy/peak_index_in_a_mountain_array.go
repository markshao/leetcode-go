package easy

// https://leetcode.com/problems/peak-index-in-a-mountain-array/
// assumption the length of A > 3
func peakIndexInMountainArray(A []int) int {
	peak := A[0]
	peakIndex := 0

	length := len(A)
	for i := 1 ; i < length; i++ {
		if A[i] > peak {
			peak = A[i]
			peakIndex = i
		} else {
			break
		}
	}

	return peakIndex
}