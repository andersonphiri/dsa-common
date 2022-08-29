// methods contains common methods
package methods

// ReverseArray reversed an array from specified start index to the specified end index
func ReverseArray[T any](a []T, startIndex int, endIndex int) {
	for startIndex < endIndex {
		a[startIndex], a[endIndex] = a[endIndex], a[startIndex]
		startIndex++
		endIndex--
	}
}

func ArrayEqual[T any](a,b []T, Compare func(one,two *T) bool) bool  {
	if a == nil || b == nil {
		return false
	}
	Na := len(a)
	Nb := len(b)
	if Na != Nb {
		return false
	}
	var i int
	for ; i <Na ; i++ {
		if !Compare(&a[i], &b[i]) {
			return false
		}
	}
	return true
}