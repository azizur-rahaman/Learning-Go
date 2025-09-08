package main

import "fmt"

func main() {
	var slice1 = []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 10, 20) // length 10, capacity 20

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println("Length of slice2:", len(slice2))
	fmt.Println("Capacity of slice2:", cap(slice2))
}
