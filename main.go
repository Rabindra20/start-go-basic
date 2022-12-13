package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
	var userName = "rab"
	// use printf and %v for placeholder
	fmt.Printf("welcome %v\n", userName) 
	// fmt.Scan(&userName)
	// fmt.Printf("%v", userName)

	// array
	var students [3]string
	fmt.Printf("students: %v \n",students)
	students[0]="ram"
	students[1]="syam"
	students[2]="gita"
	fmt.Printf("students: %v \n",students) //print all student
	fmt.Printf("students #1: %v \n",students[1]) // define student


	var identityMatrix [3][3]int = [3][3]int{ [3]int {1, 0, 0}, [3]int {0, 1, 0}, [3]int{0, 0, 1} }
	fmt.Println(identityMatrix)

	
	// array
	a := [...]int {1,2,3}
	b := &a //put a instat of &a
    b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("Length: %v\n",len(a))
} 