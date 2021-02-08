package main
import "fmt"
func main (){

	// data set or data assign
var student [3] string
student [0] = "Sourav"
student [1] = "Das"
student [2] = "Hriday"

	// data pull, data retrieve, data get
fmt.Println(student[0])
fmt.Println(len(student[0]))
fmt.Println(student[1])
fmt.Println(len(student[1]))
fmt.Println(student[2])
fmt.Println(len(student[2]))

// short hand way or string literals :

students := [3]string {"SOURAV", "DAS" , "HRIDAY"}
fmt.Println(students)

students1 := [...]string {"SD" , "DH" , "SOURAV", "DAS" , "HRIDAY"}
fmt.Println(students, len(students1))




}