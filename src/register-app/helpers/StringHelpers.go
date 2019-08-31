package helpers

// IsEmpty : check empty string of data
func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	}

	return false
}
