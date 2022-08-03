package binarysearch

func binarySearch1(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}

func binarySearch2(nums []int, target int) int {
	size := len(nums)
	l, r := 0, size-1
	for l < r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	if l != size && nums[l] == target {
		return l
	}
	return -1
}

func binarySearch3(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l+1 < r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m
		} else {
			r = m
		}
	}

	// Post-processing:
	// End Condition: left + 1 == right
	if nums[l] == target {
		return l
	}
	if nums[r] == target {
		return r
	}
	return -1
}
