package methods

import "testing"

const (
	pass = "\u2713"
	fail = "\u2717"
)

func Compare(i, j *int) bool {
	return *i == *j
}
func TestReverseArray(t *testing.T) {
	table := []struct {
		name string 
		expected []int
		input []int
		
	}{
		{name: "one", expected: []int{1,2,4,6,8,9,10}, input : []int{10,9,8,6,4,2,1}},
	}

	for _, testCase := range table {
		td := func (tf *testing.T)  {
			tf.Logf("given the need to test reverse of an array\n")
			{
				tf.Logf("when %v is reversed:\n", testCase.input)
				{
					ReverseArray(testCase.input, 0, len(testCase.input) - 1)
					if !ArrayEqual(testCase.input, testCase.expected, Compare) {
						tf.Fatalf("\t\t%s %v should be:\t%v but it was:\t%v\n",fail, testCase.input, testCase.expected, testCase.input)
					} else {
						tf.Logf("\t\t%s %v should be:\t%v\n",pass, testCase.input, testCase.expected)
					}
				}
			}
		}
		t.Run(testCase.name, td)
	}
}