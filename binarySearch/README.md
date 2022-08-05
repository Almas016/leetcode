# Когда мы можем использовать бинарный поиск?
Если мы сможем обнаружить какую-то монотонность, например, если  `condition(k) == true`, то `condition(k+1) == true`, тогда мы можем рассмотреть бинарный поиск.

# Basic template

```go
func binarySearch(nums []int, target int) int {
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
```

# Template. Left == target

```go
func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] >= target  {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
```

- [35. Search Insert Position](https://leetcode.com/problems/search-insert-position/)

# Template. Mid condition

```go
func binarySearch(n int) int {
	l, r := 1, n // could be [0, n], [1, n] etc. Depends on problem
	for l < r {
		m := l + (r-l)/2
		//if condition(mid)
		if isBadVersion(m) {
			r = m
		} else {
			l = m + 1
		}

	}
	return l // l - 1
}

func isBadVersion(x int) bool {
	return true
}
```

leetcode problems:
- [278. First Bad Version](https://leetcode.com/problems/first-bad-version/)
- [69. Sqrt(x)](https://leetcode.com/problems/sqrtx/)
- [1011. Capacity To Ship Packages Within D Days](https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/)
- [410. Split Array Largest Sum](https://leetcode.com/problems/split-array-largest-sum/)
```go
// 410. Split Array Largest Sum
func splitArray(nums []int, m int) int {
	max, sum := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	l, r := max, sum
	for l < r {
		mid := l + (r-l)/2
		if condition(mid, m, nums) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func condition(x, d int, nums []int) bool {
	count := 1
	total := 0
	for _, v := range nums {
		total += v
		if total > x {
			total = v
			count++
			if count > d {
				return false
			}
		}
	}
	return true
}
```
- [875. Koko Eating Bananas](https://leetcode.com/problems/koko-eating-bananas/)
```go
// Koko Eating Bananas
func condition(m, h int, nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += (v-1)/m + 1
	}
	return sum <= h
}
```

# Ref
- [Powerful Ultimate Binary Search Template](https://leetcode.com/discuss/general-discussion/786126/Python-Powerful-Ultimate-Binary-Search-Template.-Solved-many-problems)