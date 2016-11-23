package interview

import (
	u "github.com/dockerian/go-coding/utils"
)

// GetLongestConsecutiveIncrease returns the length of
// the longest increasing consecutive subsequence,
// and the slice, for example:
//	3 from [10, 9, 2, 5,  3, 7, 101,  18]
//	4 from [1, 2, 3, 4,  3, 5]
//	3 from [-1, 2, 3,  0]
//	2 from [-1, 0,  0, 3,  -10, 11]
//	0 from [9, 9, 9]
//	...
func GetLongestConsecutiveIncrease(arr []int) (int, []int) {
	var current, currentStart, currentEnd, saved, savedStart, savedEnd int
	var size = len(arr)

	for i := 1; i < size; i++ {
		if arr[i] > arr[i-1] {
			currentEnd = i + 1
			current++
		} else {
			if current > saved {
				saved = current
				savedStart = currentStart
				savedEnd = currentEnd
			}
			if saved >= (size - i) {
				break
			}
			current = 0
			currentStart = i
			currentEnd = i
		}
	}

	if current == 0 && saved == 0 {
		return 0, []int{}
	} else if current > saved {
		return current + 1, arr[currentStart:currentEnd]
	}

	return saved + 1, arr[savedStart:savedEnd]
}

// GetLongestIncrease returns the length of
// the longest increasing subsequence (no need to be consecutive),
// and the slice, for example:
//	4 from [10, 9, 2, 5, 3, 7, 101, 18]
//	4 from [1, 2, 3, 4]
//	5 from [1, 2, 3, 0, 9, 99]
//	9 from [-11, -10, 0, -15, -14, -12, -17, -11, 0, -9, -1, 0, 3, -10, 11]
//	0 from [-7, -7, -7]
//	...
func GetLongestIncrease(arr []int) (int, []int) {
	var size = len(arr)
	var current, saved []int
	var noIncrease = true

	// fmt.Printf("\narr(%v)=%+v\n", len(arr), arr)
	u.Debug("\narr(%v)=%+v\n", len(arr), arr)
	for m := 0; m < size-1; m++ {
		for i := m; i < size-1; i++ {
			for j := i + 1; j < size && (size-j) >= len(saved); j++ {
				var opt = arr[i]
				var previous = opt
				current = []int{previous}
				for k := j; k < size; k++ {
					if arr[k] > previous {
						opt = previous
						previous = arr[k]
						current = append(current, previous)
						noIncrease = false
					} else if arr[k] > opt {
						previous = arr[k]
						current[len(current)-1] = previous
					}
				}
				// fmt.Printf("m=%v, i=%v, j=%v, current(%v)=%+v, saved(%v)=%+v\n", m, i, j, len(current), current, len(saved), saved)
				u.Debug("m=%v, i=%v, j=%v, current(%v)=%+v, saved(%v)=%+v\n", m, i, j, len(current), current, len(saved), saved)
				if len(current) > len(saved) {
					saved = current
				}
				// never true but proves how condition should work in `for j` loop
				if len(saved) > (size - j) {
					// fmt.Println("break")
					break
				}
			}
		}
	}

	if noIncrease {
		return 0, []int{}
	} else if len(current) > len(saved) {
		return len(current), current
	}

	return len(saved), saved
}

// GetLongestSequence returns the longest subsequence in a string
func GetLongestSequence(str string, byDecending bool) string {
	if len(str) == 0 {
		return ""
	}
	start, length := 0, 1

	for i, k := 1, 1; i < len(str); i++ {
		diff := int(str[i]) - int(str[i-1])
		sequential := byDecending && diff == -1 || !byDecending && diff == 1

		if sequential {
			// u.Debug("i = %v, k+1 = %v [start= %v, length= %v] in '%v'\n",
			// 	i, k+1, start, length, str)
			if k++; k > length {
				start = i + 1 - k
				length = k
			}
			continue
		}
		k = 1
	}

	return str[start : start+length]
}