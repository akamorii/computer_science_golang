package main

import (
	"fmt"
	se "main/algoritms"
	ll "main/datastruct"
)

func main() {
	test_linked_list()
	test_binary_search()
}

func test_binary_search() {
	zxc := [13]int{1, 4, 7, 90, 128, 133, 145, 147, 150, 151, 154, 156, 157}

	fmt.Println(se.Bi_search(zxc, 133))
}

func test_linked_list() {
	arr := ll.Linked_list{}
	arr.Append(14)
	arr.Append(15)
	arr.Append(16)
	// fmt.Println(arr)
	zxc := arr.Take_element(0)
	fmt.Println(zxc)
	arr.Show_arr()
}
