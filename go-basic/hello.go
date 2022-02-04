package main

import (
	"fmt"
	"go-basic/calculation"
)

func main(){
	fmt.Println("Hello, World!")
	result := calculation.Add(1,1)
	fmt.Println(result)
}