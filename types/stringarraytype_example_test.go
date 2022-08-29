package types

import "fmt"

func ExampleTestStringArray() {
	a := []int{1,2,3,4,5,6,7,8}
	fmt.Println(StringArrayType[int](a).String())
	// Output:
	// 12345678
}