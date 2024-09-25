package supportfiles

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func FindMaxLength2DArray(arr [][][]string) [][]string {
	var maxLength int
	var max2DArray [][]string

	// Loop through the 3D array
	for _, twoDArr := range arr {
		// Check if the current 2D array has a greater length than the current maximum
		if len(twoDArr) > maxLength {
			maxLength = len(twoDArr)
			max2DArray = twoDArr
		}
	}
	return max2DArray
}
