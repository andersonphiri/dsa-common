package bitwise

import (
	"fmt"
	"math/bits"
	"testing"
)
var upTo uint8 = 30
const (
	pass = "\u2713"
	fail = "\u2717"
)

func init() {
	// Initiliaze()
}
func TestPrintLookupTable(t *testing.T) {
	PrintLookupTable(upTo);
}

func TestPrintLookupTableValue(t *testing.T) {
	var testValue uint8 = 203;
	PrintSnapLookupTable(testValue)
	fmt.Println();
}

func TestCountSetBits(t *testing.T) {
	table := []struct {
		input uint
		expected int
		name string
	}{
		{input: 0xFF, expected: 8, name: "0xFF"},
		{input: 0xFFAE, expected: bits.OnesCount(0xFFAE), name: "0xFFAE"},
		{input: 0xFFAE, expected: bits.OnesCount(0xFFAE), name: "0xFFAE"},
	}
	for _, test := range table {
		tf := func (tf *testing.T)  {
			value := CountSetBits(test.input)
			if value != test.expected {
				tf.Fatalf("\t%s\t\t%d has a count of set bit of:\t%d but it got:\t%d\n", fail, test.input, test.expected, value )
			} else {
				tf.Logf("\t%s\t\t%d has a count of set bit of:\t%d\n", pass, test.input,  value )
			}
		}
		t.Run(test.name, tf)
	}
	fmt.Println();
}