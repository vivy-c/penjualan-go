package main 

import (
	"fmt"
)

func main() {
	size := 0  
	fmt.Print("n+m=")  
	fmt.Scanln(&size)  
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