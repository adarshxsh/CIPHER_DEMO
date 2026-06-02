package main

import "fmt"

func main(){
	var array [5]int
	fmt.Print("Enter the 5 elements of the array: ")
	for i := 0;i<5;i++ {
		fmt.Scan(&array[i])
	}
	array = sort(array)
	sl_idx := sl(array)
	if sl_idx != -1 {
		fmt.Print("The Second largest is: ")
		fmt.Print(array[sl_idx])
	} else {
		fmt.Println("Second Largest Element doesn't Exist")
	}
}

func sort(arr [5]int) [5]int {
	for i := 0;i<5;i++{
		min := arr[i]
		idx_min := i
		for j:=i;j<5;j++{
			if(arr[j]<min){
				min = arr[j]
				idx_min = j
			}
		}
		arr[idx_min] = arr[i]
		arr[i] = min
	}
	return arr
}

func sl (a [5]int) int {
	y := 3
	for (a[y+1] == a[y] && y != 0){
		y--
	}
	if(a[y+1] != a[y]){
		return y
	} else {
		return -1
	}
}
