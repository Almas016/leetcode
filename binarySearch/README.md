# Template 1
<details><summary>Code</summary>

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

</details>


# Template. l==target
leetcode problems:
- [Find K Closest Elements](https://leetcode.com/problems/find-k-closest-elements/)

<details><summary>Code</summary>

```go
l, r := 0, len(nums)-1
for l < r {
    mid := (l + r)/2
    if nums[mid] < target {
        l = mid + 1
    } else {
        r = mid
    }
}
```

</details>

# Template 3
Он используется для поиска элемента или условия, требующего доступа к текущему индексу и его ближайшим левым и правым соседним индексам в массиве.

<details><summary>Code</summary>

```go
func binarySearch(nums []int, target int) int {
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
```

</details>