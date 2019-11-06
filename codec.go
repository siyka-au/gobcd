package bcd

// Encode Encode uint64 to a BCD byte array
func Encode(x uint64) []byte {
	if x == 0 {
		return []byte{0x0f}
	}
	var n int
	for xx := x; xx > 0; n++ {
		xx = xx / 10
	}
	bcd := make([]byte, (n+1)/2)
	if n%2 == 1 {
		hi, lo := byte(x%10), byte(0x0f)
		bcd[(n-1)/2] = hi<<4 | lo
		x = x / 10
		n--
	}
	for i := n/2 - 1; i >= 0; i-- {
		hi, lo := byte((x/10)%10), byte(x%10)
		bcd[i] = hi<<4 | lo
		x = x / 100
	}
	return bcd
}

func timesTenPlusCatchingOverflow(x uint64, digit uint64) (uint64, error) {
	x5 := x<<2 + x
	if int64(x5) < 0 || x5<<1 > ^digit {
		return 0, OverflowError{}
	}
	return x5<<1 + digit, nil
}

// Decode Decode a BCD byte array to uint64
func Decode(bcd []byte) (x uint64, err error) {
	for i, b := range bcd {
		hi, lo := uint64(b>>4), uint64(b&0x0f)
		if hi > 9 {
			return 0, BadDigitError{"hi", hi}
		}
		x, err = timesTenPlusCatchingOverflow(x, hi)
		if err != nil {
			return 0, err
		}
		if lo == 0x0f && i == len(bcd)-1 {
			return x, nil
		}
		if lo > 9 {
			return 0, BadDigitError{"lo", lo}
		}
		x, err = timesTenPlusCatchingOverflow(x, lo)
		if err != nil {
			return 0, err
		}
	}
	return x, nil
}
