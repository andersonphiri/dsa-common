package mathematics

import ( 
	"testing"
	"github.com/andersonphiri/ds-common/types"
)

func TestFactorialLarge(t *testing.T) {
	table := []struct {
		expected string
		actual   string
		n        uint
		name     string
		skipCheck bool
	}{
		{expected: "5040", n: 7, name: "Testing7"},
		{expected: "6402373705728000", n: 18, name: "Testing18"},
		{expected: "15511210043330985984000000", n: 25, name: "Testing25"},
		{expected: "93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000", n: 100, name: "Testing100", skipCheck: true},
	}
	for _, testcase := range table {
		td := func(t *testing.T) {
			res := ComputeMongoFactorial(testcase.n)
			testcase.actual = types.StringArrayType[int](res).String()
			if testcase.skipCheck {
				t.Logf("%s\t FactorialLarge(%d) was\t\t:%s, test skipped", pass,testcase.n, testcase.actual)
				return
			}
			if testcase.expected != testcase.actual {
				t.Fatalf("%s\texpected\tFactorialLarge(%d)\tshould be:\t\t%s but it was:\t\t%s\n", fail,
					testcase.n, testcase.expected, testcase.actual)
			} else {
				t.Logf("%s\t is:\t\t:%s", pass, testcase.actual)
			}
		}
		t.Run(testcase.name, td)
	}

}
