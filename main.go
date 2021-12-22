package main 

import (
	"fmt"
)

func main() {
	
	n := 0
	m := 0
	fmt.Print("n=")  
	fmt.Scanln(&n)  
	fmt.Print("m=")  
	fmt.Scanln(&m) 
	fmt.Print("n+m=")  
	size := n+m
	fmt.Println(size)  
	fmt.Println("isi elemen")  
	elements := make([]int, size)  
	for i := 0; i < size; i++ {  
	 fmt.Scanln(&elements[i])  
	}  
	fmt.Println("elemen yang diinput:", elements)  
	result := 0  
	 
	for i := 0; i < size; i++ {  
	 result += elements[i]  
	 
	}  
	fmt.Println("hasil:", result)  
}