package util

import "golang.org/x/exp/constraints"

// Abs returns the absolute value of x. If x is negative, it returns -x.
func Abs[I constraints.Integer | constraints.Float](x I) I {
	if x < 0 {
		return -x
	}
	return x
}
