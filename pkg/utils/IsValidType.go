package utils

func IsValidType(x interface{}) bool {
	switch x.(type) {
	case string,
		int,
		float64,
		byte,
		bool,
		map[string]string,
		map[string]int,
		map[string]bool,
		[]string,
		[]int:
		return true
	default:
		return false
	}
}
