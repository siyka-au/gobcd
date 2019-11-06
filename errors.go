package bcd

import "fmt"

// BadDigitError An error when a bad BCD digit is encoountered
type BadDigitError struct {
	v   string
	val uint64
}

func (e BadDigitError) Error() string {
	return fmt.Sprintf("Bad digit in BCD decoding: %s = %d", e.v, e.val)
}

// OverflowError An error when an overflow occurs when generating a BCD value
type OverflowError struct{}

func (e OverflowError) Error() string {
	return "Overflow occurred in BCD decoding"
}
