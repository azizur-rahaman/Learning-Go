package main

import "fmt"

func main() {
	var array_name [5]int = [5]int{1, 2, 3, 4, 5}
	var array2 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr1 := [3]int{1: 5, 2: 10} // index 0 will be 0, index 1 will be 5, index 2 will be 10

	fmt.Println(array_name)
	fmt.Println(array2)
	fmt.Println(arr1)

	fmt.Println("Length of array2:", len(array2))
}
