package set

import "golang.org/x/exp/constraints"

func SortedIntersectionIndex[T, U any](lhs []T, rhs []U, compare func(lhs T, rhs U) int) []int {
	lhsIndex, lhsLength := 0, len(lhs)
	rhsIndex, rhsLength := 0, len(rhs)
	indexes := make([]int, 0, lhsLength)

	for lhsIndex < lhsLength && rhsIndex < rhsLength {
		c := compare(lhs[lhsIndex], rhs[rhsIndex])
		switch {
		case c < 0:
			lhsIndex++
		case c > 0:
			rhsIndex++
		default:
			indexes = append(indexes, lhsIndex)
			lhsIndex++
			rhsIndex++
		}
	}

	return indexes
}

func SortedOrderedIntersectionIndex[T constraints.Ordered](lhs, rhs []T) []int {
	return SortedIntersectionIndex(lhs, rhs, OrderedCompare[T])
}

func SortedIntersection[T, U any](lhs []T, rhs []U, compare func(lhs T, rhs U) int) []T {
	return indexToValue(SortedIntersectionIndex(lhs, rhs, compare), lhs)
}

func SortedOrderedIntersection[T constraints.Ordered](lhs, rhs []T) []T {
	return SortedIntersection(lhs, rhs, OrderedCompare[T])
}

func SortedDifferenceIndex[T, U any](lhs []T, rhs []U, compare func(lhs T, rhs U) int) []int {
	intersections := SortedIntersectionIndex(lhs, rhs, compare)
	intersectionsIndex, intersectionsLength := 0, len(intersections)
	lhsIndex, lhsLength := 0, len(lhs)
	differences := make([]int, 0, lhsLength-intersectionsLength)

	for lhsIndex < lhsLength && intersectionsIndex < intersectionsLength {
		if intersections[intersectionsIndex] != lhsIndex {
			differences = append(differences, lhsIndex)
		} else {
			intersectionsIndex++
		}
		lhsIndex++
	}

	for lhsIndex < lhsLength {
		differences = append(differences, lhsIndex)
		lhsIndex++
	}

	return differences
}

func SortedOrderedDifferenceIndex[T constraints.Ordered](lhs, rhs []T) []int {
	return SortedDifferenceIndex(lhs, rhs, OrderedCompare[T])
}

func SortedDifference[T, U any](lhs []T, rhs []U, compare func(lhs T, rhs U) int) []T {
	return indexToValue(SortedDifferenceIndex(lhs, rhs, compare), lhs)
}

func SortedOrderedDifference[T constraints.Ordered](lhs, rhs []T) []T {
	return SortedDifference(lhs, rhs, OrderedCompare[T])
}

func indexToValue[T any](index []int, value []T) []T {
	set := make([]T, len(index))
	for i := range set {
		set[i] = value[index[i]]
	}
	return set
}
