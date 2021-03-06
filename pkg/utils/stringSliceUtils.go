package utils

// AddStringIfMissing adds a string to a slice of strings
// The provided slice is not changed
// returns true, slice with appended string if string was not in the list
// returns false, old list if string was already in the list
func AddStringIfMissing(slice []string, s string) (bool, []string) {
	for _, item := range slice {
		if item == s {
			return false, slice
		}
	}
	return true, append(slice, s)
}

// RemoveString removes a string from a slice of strings
// The provided slice is not changed
// returns true, slice with removed string if string was contained in the list
// returns false, old list if string was not in the list
func RemoveString(slice []string, s string) (bool, []string) {
	removed := false
	result := []string{}
	for _, item := range slice {
		if item == s {
			removed = true
			continue
		}
		result = append(result, item)
	}
	return removed, result
}

// StringSliceContains determines whether the given string slice
// contains a given element.
func StringSliceContains(slice []string, elem string) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}
