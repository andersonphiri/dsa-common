package bitwise

import "fmt"
var LookupTable [256]uint8
var LookupMax uint8 = 255

func init() {
	
	var i uint8 = 1
	for ; i < LookupMax ; i++ {
		LookupTable[i] = (i & 1) + LookupTable[i / 2] // loop up to 254 to evade overflow
	}
	LookupTable[LookupMax] = 8

}

func PrintLookupTable(upTo uint8) {
	var i uint8 = 0;
	fmt.Printf("N\t\tBINARY\t\tSET_BITS_COUNT\n")
	for ; i <= upTo ; i++ {

		fmt.Printf("%d\t\t%b\t\t%5d\n", i, i, LookupTable[i & LookupMax])
	} 
	fmt.Println()
}

func PrintSnapLookupTable(value uint8) {
	fmt.Printf("N\t\tBINARY\t\tSET_BITS_COUNT\n")
	fmt.Printf("%d\t\t%b\t\t%5d\n", value, value, LookupTable[value & LookupMax])
	fmt.Println()
}

func GetSnapLookupTable(value uint8) uint8 {
	return LookupTable[value & LookupMax];
}

// CountSetBitsInt counts the number of set bits
func CountSetBits(value uint) (result int) {
	var lookupMask uint = 0xFF
	 result  = int(LookupTable[value & lookupMask])
	 value >>= 8
	  for ; value > 0 ; {
		result += int(LookupTable[value & lookupMask])

		value >>= 8
	  }
	  return

}