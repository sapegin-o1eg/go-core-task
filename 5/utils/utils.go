package utils

func Intersection(a, b []int) (ok bool, result []int) {
	if len(a) == 0 || len(b) == 0 {
		return false, []int{}
	}

	set, check := a, b
	if len(a) > len(b) {
		set, check = check, set
	}

	m := make(map[int]struct{}, len(set))
	for _, v := range set {
		m[v] = struct{}{}
	}

	result = make([]int, 0, min(len(set), len(check)))
	for _, v := range check {
		if _, exists := m[v]; exists {
			result = append(result, v)
			delete(m, v)
		}
	}

	return len(result) > 0, result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
