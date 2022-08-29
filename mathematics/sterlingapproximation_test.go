package mathematics
import (
	"testing"
)

const (
	pass = "\u2713"
	fail = "\u2717"
)

func TestSterlingApproximation(t *testing.T) {
	var table = []struct{
		Name string
		Input int
		Output int64 
		OutputRound int64
		SkipCheck bool
	}{
		{Name: "20", Input: 20 , Output: 19, OutputRound: 19},
		{Name: "10Pow9", Input: 1_000_000_000 , Output: 8_565_705_527 , OutputRound: 8_566_000_000},
		// {Name: "20", Input: 20 , Output: 19},
		// {Name: "20", Input: 20 , Output: 19},
		{Name: "10Pow18", Input: 100000_000_000_000_000 , Output: 1656570551809674752 , OutputRound: -1, SkipCheck: true},
	}
	for _, input := range table {
		tfunc := func (tf *testing.T)  {
			// given
			t.Logf("\tgiven the need to compute the number of digits for %d!\n", input.Input)
			{
				// when 
				t.Logf("when %d! is computed using sterling approximation\n", input.Input) 
				{
					// should 
					result := SterlingApproximation(input.Input)
					if input.SkipCheck {
						tf.Logf("\t%s Test: %s, %d! should be %d digits and it got %d, expected rounded to %d (test to be igonred)\n", pass,input.Name, input.Input, input.Output, result, input.OutputRound)
						return
					}
					if result != input.Output  {
						tf.Fatalf("\t%s Test: %s, %d! should be %d digits but it got %d, expected rounded to %d\n", fail,input.Name, input.Input, input.Output, result, input.OutputRound)
					} else {
						tf.Logf("\t%s Test: %s, %d! should be %d digits and it got %d, expected rounded to %d\n", pass,input.Name, input.Input, input.Output, result, input.OutputRound)
					}
				}
			}
		}
		t.Run(input.Name,tfunc)
	}
}


func TestSterlingApproximationInt(t *testing.T) {
	var table = []struct{
		Name string
		Input int
		Output int
		OutputRound int64
		SkipCheck bool
	}{
		{Name: "20", Input: 20 , Output: 19, OutputRound: 19},
		{Name: "10Pow9", Input: 1_000_000_000 , Output: 8_565_705_527 , OutputRound: 8_600_000_000, SkipCheck: true},
		{Name: "10Pow18", Input: 1000_000_000_000_000_000 , Output: 1656570551809674752 , OutputRound: -1, SkipCheck: true},
	}
	for _, input := range table {
		tfunc := func (tf *testing.T)  {
			// given
			t.Logf("\tgiven the need to compute the number of digits for %d!\n", input.Input)
			{
				// when 
				t.Logf("when %d! is computed using sterling approximation\n", input.Input) 
				{
					// should 
					result := SterlingApproximation(input.Input)
					if input.SkipCheck {
						t.Logf("\t%s Test: %s, %d! should be %d digits and it got %d, expected rounded to %d (test to be igonred)\n", 
						pass,input.Name, input.Input, input.Output, result, input.OutputRound)
						return
					}
					if result != input.OutputRound  {
						t.Fatalf("\t%s Test: %s, %d! should be %d digits but it got %d, expected rounded to %d\n", 
						fail,input.Name, input.Input, input.Output, result, input.OutputRound)
					} else {
						tf.Logf("\t%s Test: %s, %d! should be %d digits and it got %d, expected rounded to %d\n", 
						pass,input.Name, input.Input, input.Output, result, input.OutputRound)
					}
				}
			}
		}
		t.Run(input.Name,tfunc)
	}
}
