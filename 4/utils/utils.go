package utils

func Difference(source, exclude []string) []string {
	switch {
	case source == nil:
		return []string{}
	case exclude == nil:
		return source
	case len(source) == 0:
		return []string{}
	case len(exclude) == 0:
		return source
	}

	excludeSet := make(map[string]bool)
	resultSlice := make([]string, 0)

	for _, item := range exclude {
		excludeSet[item] = true
	}

	for _, item := range source {
		if _, ok := excludeSet[item]; !ok {
			resultSlice = append(resultSlice, item)
		}
	}

	return resultSlice
}
