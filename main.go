package main

import "fmt"
import "reflect"

func main (){

// slice

//var student [3]string
//student[0] = "Sourav"
//student[1] = "Das"
//student[2] = "Hriday"
//x :=student[0:2]

	//short hand format 
//x := make([]string,0) 
//fmt.Println(x)
	
	//long hand format
var fruits []string
fruits= append(fruits,"Apple", "Banana" , "Mango")
//fmt.Println(fruits,len(fruits))
	
//fmt.Printf("%T\n",fruits) //for to see variable types
//fmt.Printf("%T\n",student)

b := reflect.TypeOf(fruits).kind(),String()
fmt.Println(b)
 
}