package connectip

import (
	"golang.org/x/exp/constraints"
)

// isNaN reports whether x is a NaN without requiring the math package.
// This will always return false if T is not floating-point.
func isNaN[T constraints.Ordered](x T) bool {
	return x != x
}

func min[T constraints.Ordered](x, y T) T {
	if isNaN(x) {
		return x
	}
	if isNaN(y) {
		return y
	}
	if x < y {
		return x
	}
	return y
}

func max[T constraints.Ordered](x, y T) T {
	if isNaN(x) {
		return x
	}
	if isNaN(y) {
		return y
	}
	if x < y {
		return y
	}
	return x
}
