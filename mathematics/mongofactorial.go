package mathematics

import "github.com/andersonphiri/ds-common/methods"

const marginOfError = 10

// ComputeMongoFactorial computes the factorial of a huge number and store the result as an array in reverse order
// to get the correct value, reverse t
func ComputeMongoFactorial(n uint) (result []int) {
	expectedN := SterlingApproximationInt(int(n)) // expected number of digits
	result = make([]int, 0, expectedN + marginOfError ) // to minimize the cost of copying array items when they full
	result = append(result, 1)
	var i int = 2
	for ; i <= int(n) ; i++ {
		var carry int
		for j := 0; j < len(result); j++ {
			val := i * result[j] + carry
			result[j] = val % 10
			carry = val/10
		}

		for carry != 0 {
			result = append(result, carry % 10)
			carry = carry / 10
		}
	}
	// 

	methods.ReverseArray(result, 0, len(result) -1)
	return
	
}