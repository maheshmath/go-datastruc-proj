package main

import "fmt"

func arrByVal(x [5]int) {
	fmt.Println(x[0])
	x[0] = 50
	fmt.Println(x[0])
}

func arrByPointer(x *[5]int) {
	fmt.Println(x[0])
	x[0] = 100
	fmt.Println(x[0])
}

func modifyArr(x []int) {
	fmt.Println(x[0])
	x[0] = 10
	fmt.Println(x[0])
}
func main() {
	fmt.Println("Hello, Array Example!")
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	arrByVal(arr)
	fmt.Println(arr)

	arrByPointer(&arr)
	fmt.Println(arr)

	modifyArr(arr[:])
	fmt.Println(arr)
}
