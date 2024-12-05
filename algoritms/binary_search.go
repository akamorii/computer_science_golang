package algoritms

func Bi_search(arr []int, find int) bool {
	var middle int
	first := 0
	last := len(arr) - 1
	for first <= last {
		middle = first + (last-first)/2
		// fmt.Println(middle)
		if arr[middle] == find {
			return true
		} else if arr[middle] > find {
			last = middle - 1
		} else {
			first = middle + 1
		}

	}
	return false
}
