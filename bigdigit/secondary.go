package bigdigit

func RemoveBeginZero(arr []int) []int {
	lastIndex := len(arr) - 1
	for ; lastIndex > 0 && arr[lastIndex] == 0; lastIndex-- {
	}
	return arr[:lastIndex+1]
}
