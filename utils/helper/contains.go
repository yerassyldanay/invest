package helper

/*
	this function has been created to check whether a slice contains an element
*/
func DoesASliceContainElement(aslice []string, target string) bool {
	for _, value := range aslice {
		if value == target {
			return true
		}
	}

	return false
}
