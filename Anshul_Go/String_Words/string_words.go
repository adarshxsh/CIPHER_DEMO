package main

import "fmt"

func main(){
	count := 0
	fmt.Print("Enter the string ")
	var str string
	fmt.Scanf("%s",&str)
	fmt.Println(str)
	for range str{
		count++;
	}
	fmt.Print(count)
	// fmt.Print(len(str))  --> direct method
	}