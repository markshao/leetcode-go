package easy

import "testing"
import "fmt"

func TestFlippingImageTest(*testing.T) {
	var inputData = [][]int{
		{1, 1, 0},
		{1, 0, 1},
		{0, 0, 0},
	}
	result := FlipAndInvertImage(inputData)
	fmt.Println(result)
}
