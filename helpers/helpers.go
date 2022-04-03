package helpers

func InArray(x int, arr []int) bool {
	for _, v := range arr {
		if x == v {
			return true
		}
	}
	return false
}
