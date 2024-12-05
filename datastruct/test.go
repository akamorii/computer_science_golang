package main

import (
	ll "datastruct/linked_list"
	"fmt"
)

func main() {
	test_linked_list()

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
