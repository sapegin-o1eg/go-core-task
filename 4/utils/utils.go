package utils

func Difference(source, exclude []string, args ...bool) []string {
	if len(source) == 0 {
		return []string{}
	}

	if len(exclude) == 0 {
		return source
	}

	removeDuplicates := false
	if len(args) > 0 && args[0] {
		removeDuplicates = args[0]
	}

	excludeSet := make(map[string]bool)
	resultSlice := make([]string, 0)
	resultSet := make(map[string]bool, 0)

	for _, item := range exclude {
		excludeSet[item] = true
	}

	for _, item := range source {
		if _, ok := excludeSet[item]; !ok {
			if removeDuplicates {
				if _, ok := resultSet[item]; !ok {
					resultSet[item] = true
					resultSlice = append(resultSlice, item)
				}
			} else {
				resultSlice = append(resultSlice, item)
			}
		}
	}

	return resultSlice
}
