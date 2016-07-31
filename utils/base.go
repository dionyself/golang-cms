package utils

/*
func Containss(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
*/

// Contains Verify if slice contains x string
func Contains(stringSlice []string, stringToSearch string) bool {
	for _, stringElement := range stringSlice {
		if stringElement == stringToSearch {
			return true
		}
	}
	return false
}
