package utils

func RemoveBeginZero(arr []int) []int {
	lastIndex := len(arr) - 1
	for ; lastIndex > 0 && arr[lastIndex] == 0; lastIndex-- {
	}
	return arr[:lastIndex+1]
}

func CompareModule(d1 []int, d2 []int) int8 {
	if len(d1) != len(d2) {
		if len(d1) < len(d2) {
			return -1
		}
		return 1
	}
	for i := len(d1) - 1; i >= 0; i-- {
		if d1[i] < d2[i] {
			return -1
		} else if d1[i] > d2[i] {
			return 1
		}
	}
	return 0
}
