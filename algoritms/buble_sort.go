package algoritms

func Buble_sort(arr *[]int) {
	arr_length := len(*arr) - 1
	for i := 0; i < arr_length; i++ {
		for j := 0; j < arr_length-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}

		}
	}
}
