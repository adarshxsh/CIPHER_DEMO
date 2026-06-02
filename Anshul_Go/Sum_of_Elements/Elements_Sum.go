package main

import "fmt"

func main(){
	var array [5]int
	fmt.Print("Enter the 5 elements of the array: ")
	for i := 0;i<5;i++ {
		fmt.Scan(&array[i])
	}
	fmt.Print(sum(array))
}

func sum(arr [5]int) int {
	var sum = 0
	for _,val := range arr{
		sum+=val
	}
	return sum

}