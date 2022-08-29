package mathematics


import (
	"math"
)

// SterlingApproximation computes the approximate number of digits of a factorial of n, for example if n = 20,
// the expected number of digits for 20! is, factorial(20) has 19 digits
// the formula: given a simple number say 235, number of digits = 1 + Floor(math.log(235) base 10) or Ceil (math.log(235) base 10) = 1 = Floor(2.371...) or Ceil(2.371..)
// = 3
// formula: log(n! base 10) = ( N*ln(N) - N + ln(2*pi*N)  ) / ln10
func SterlingApproximation(n int) int64 {
	ln10 := math.Log(10)
	floatN := float64(n)
	nLnN := floatN * math.Log(float64(n))
	pi := float64(22) / float64(7)
	ln2piN := math.Log( float64(2*n) * pi )
	result := ((nLnN)/ln10) - (floatN/ln10) + (ln2piN/ln10) 
	return int64(result)

}

// SterlingApproximationInt computes the approximate number of digits of a factorial of n, for example if n = 20,
// the expected number of digits for 20! is, factorial(20) has 19 digits
// the formula: given a simple number say 235, number of digits = 1 + Floor(math.log(235) base 10) or Ceil (math.log(235) base 10) = 1 = Floor(2.371...) or Ceil(2.371..)
// = 3
// formula: log(n! base 10) = ( N*ln(N) - N + ln(2*pi*N)  ) / ln10
func SterlingApproximationInt(n int) int {
	ln10 := math.Log(10)
	floatN := float64(n)
	nLnN := floatN * math.Log(float64(n))
	pi := float64(22) / float64(7)
	ln2piN := math.Log( float64(2*n) * pi )
	result := ((nLnN)/ln10) - (floatN/ln10) + (ln2piN/ln10) 
	return int(result)

}