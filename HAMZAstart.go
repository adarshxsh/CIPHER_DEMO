package main

import "fmt"

func main() {
	//Q1 Minimum of Numbers in Array
	arr:=[]int{5,6,12,2,1,67}
	min:=arr[0]
	for _,val := range arr{
		if val<min{
		val=min
	}
	fmt.Println(min)
}

//Q2SUM of elemnts
    sum:=0
    for _,val := range arr{
	sum+=val
}
    fmt.Println(sum)
//Q3COUNT OCCURENCES
    d:=0
    var n int

    fmt.Scan(&n)
	for _,val := range arr{
		if val==n{
			d++;
		}
	}
	fmt.Println(d)
//Q4 CHECK PALLINDROME
   var s string
   fmt.Scan(&s)
   isPall:=true
   for i := 0; i < len(s)/2; i++{
	if s[i]!=s[len(s)-1-i]{
		isPall=false
		break;
	}
   }
   fmt.Println(isPall)
//Q5SUM Of FIRST N NUMBERS
   var j int;
   fmt.Scan(&j)
   uSUM:=0
   for i := 1; i <= j; i++{
	uSUM+=i
   }
   fmt.Println(uSUM)



  



}
