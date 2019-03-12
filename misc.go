package n

// MergeMap b into a and returns the new modified a
// b takes higher precedence and will override a
func MergeMap(a, b map[string]interface{}) map[string]interface{} {
	switch {
	case (a == nil || len(a) == 0) && (b == nil || len(b) == 0):
		return map[string]interface{}{}
	case a == nil || len(a) == 0:
		return b
	case b == nil || len(b) == 0:
		return a
	}

	for k, bv := range b {
		if av, exists := a[k]; !exists {
			a[k] = bv
		} else if bc, ok := bv.(map[string]interface{}); ok {
			if ac, ok := av.(map[string]interface{}); ok {
				a[k] = MergeMap(ac, bc)
			} else {
				a[k] = bv
			}
		} else {
			a[k] = bv
		}
	}

	return a
}

// Range creates slice of the given range of numbers inclusive
func Range(min, max int) []int {
	result := make([]int, max-min+1)
	for i := range result {
		result[i] = min + i
	}
	return result
}

// SetValueOrDefault sets the value of the given string 'target' to the
// given 'value' if value is not empty else 'defaulty' and returns the value as well
func SetValueOrDefault(target *string, value, defaulty string) string {
	if value != "" {
		*target = value
	} else {
		*target = defaulty
	}

	return *target
}

// SetValueIfEmpty sets the value of the given string to the other if not empty
// and returns the target as well.
func SetValueIfEmpty(target *string, value string) string {
	if *target == "" {
		*target = value
	}
	return *target
}

// ValueOrDefault returns the value or defaulty if the value is empty
func ValueOrDefault(value, defaulty string) string {
	if value != "" {
		return value
	}
	return defaulty
}
