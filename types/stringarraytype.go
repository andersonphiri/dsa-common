// types contains all reusable types
package types

import (
	"fmt"
	"strings"
)
// StringArrayType create a string representation of the array of type T
// the string format uses the value itself
type StringArrayType[T any] []T


// String create a string representation of the array of type T
func (sr StringArrayType[T]) String() string {
	var sb strings.Builder
	for _, v := range sr {
		sb.WriteString(fmt.Sprintf("%v", v))
	}
	return sb.String()
}
