package utils

/*
	this function has been created to check whether a slice contains an element
 */
func Does_a_slice_contain_element(aslice []interface{}, target interface{}) bool {
	for _, value := range aslice {
		if value == target {
			return true
		}
	}

	return false
}
