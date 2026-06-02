package main

import "fmt"

func main() {
	var array = [5]int{1,2,3,4,5}

	var newArr [5]int
	newArr = copyArr(array)
	fmt.Println(newArr)
}

func copyArr(arr [5]int) [5]int {
	var newArray [5]int 
	for idx,item := range arr{
		newArray[idx] = item
	}
	// Alternate copy(newArray,arr)
	return newArray
}