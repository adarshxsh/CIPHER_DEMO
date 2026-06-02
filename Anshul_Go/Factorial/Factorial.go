package main

import "fmt"

func main(){
	var num int
	fmt.Print("Enter the Number: ")
	fmt.Scan(&num)
	fmt.Print(factorial(num))
	}

func factorial(n int) int {
	if(n==0){
		return 1
	} else{
		return n*factorial(n-1)
	}
}