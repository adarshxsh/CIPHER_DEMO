package main

import "fmt"
func swap(x, y int) {
	fmt.Printf("Before swap: x = %d, y = %d\n", x, y)
	x, y = y, x
	fmt.Printf("After swap: x = %d, y = %d\n", x, y)
	fmt.Println()
}
func printsmallestlargest(arr [] int) {
	if len(arr) == 0 {
		fmt.Println("Array is empty")
		return
	}
	var smallest int =  arr[0]
	var largest int =  arr[0]
	for _, value := range arr {
		if value < smallest {
			smallest = value
		}
		if value > largest {
			largest = value
		}
	}
	fmt.Printf("The smallest: %d, and the largest: %d\n", smallest, largest)
	fmt.Println()
}

func checksorted(arr []int) bool {

	var asc bool = true
	var desc bool = true


	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] { 
			asc = false
		}
		if arr[i] > arr[i-1] {
			desc = false
		}
	}
	return  asc || desc
}
func checkAnagram(str1 string,str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	var count1[26]int
	var count2[26]int
	for i := 0; i < len(str1); i++ {
		count1[str1[i]-'a']++
		count2[str2[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if count1[i] != count2[i] {
			return false
		}
	}
	return true
}
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
	
func main() {

	fmt.Println("Hello World")
	var a, b int = 10,20 
	swap(a,b)
	var arr = []int {6,7,5,8,9,4,3}
	fmt.Println(arr)
	printsmallestlargest(arr)
	fmt.Println("The array is sorted:", checksorted(arr))
	var str1 = "earth"
	var str2 = "heart"
	fmt.Println("The strings,", str1, "and", str2, "are anagrams:", checkAnagram(str1, str2))

	fmt.Println("fibonacci value of 10:", fibonacci(10))


}
