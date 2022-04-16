package set

import "golang.org/x/exp/constraints"

func OrderedCompare[T constraints.Ordered](lhs, rhs T) int {
	if lhs < rhs {
		return -1
	}
	if lhs > rhs {
		return 1
	}
	return 0
}
