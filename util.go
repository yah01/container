package container

import "golang.org/x/exp/constraints"

// Search for the first element which is not less than target
func Search[S ~[]T, T constraints.Ordered](slice S, target T) int {
	if len(slice) == 0 {
		return 0
	}

	var (
		l   = 0
		r   = len(slice) - 1
		mid int
	)

	for l < r {
		mid = (l + r) >> 1
		if slice[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if slice[r] < target {
		return -1
	}

	return len(slice)
}
