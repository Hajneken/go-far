package main

// every go program must start with line defining a package

/*
Run go program using: 		go run syntaxCheatsheet.go
Compile go program using: 	go build syntaxCheatsheet.go
Init go module using: 		go mod init example.com/username/repo
*/

// importing go packages can be done by listing desired packages in the import statement
// note that no "," is used to deliminate packages
// "fmt" stands for format
import (
	"fmt"
	"log"
	"math"
	"time"
)

// alias for fmt.Println
var print = fmt.Println

// variables outside of functions are global variables
var globalGreeting string = "Hello World"

// main function is the entry point for the program
func main() {
	// print to console, if you like typing use fmt.Println, otherwise use print alias we declared above
	fmt.Println("Hello World")

	/* VARIABLES */

	// declare variable like this
	var localVar string
	// assign it
	localVar = "My Local String"
	print(localVar)
	// or just do it in one take
	var oneLinerVar int = 1
	print("This has been declared as one liner: ", oneLinerVar)
	// or using shorthand for declaring and initializing i.e. instead of `var otherLocalVar string = "My Other Local String"`
	otherLocalVar := "My Other Local String"
	print(otherLocalVar)
	// multiple assignment is supported
	var myInt1, myInt2 int = 1, 2
	print("Declared 2 ints", myInt1, "and", myInt2)
	// automatic inference is supported
	var mysteriousVar = 0.3
	print("Go can infer variable types automatically", mysteriousVar)
	// constants are supported
	const pi = 3.14159265359
	// with const you can do arithmetic with arbitrary precision
	print("Pi is", pi, "and 2*pi is", 2*pi)

	/* DATA TYPES */

	// structs are used instead of classes, they have the same purpose
	// but inheretance is not supported
	type name struct {
		firstName  string
		middleName string
		lastName   string
	}

	// go embraces composition over inheritance, i.e. put types together instead of inheriting from them
	// capital letters are used to denote public fields, i.e. Person is exported and can be used elsewhere
	type Person struct {
		// structs can be embedded in other structs
		name name
		age  int
	}

	// initialize object of type person
	john := Person{name{"John", "Doe", "Smith"}, 42}
	// the default toString() method is sensible
	print("We consider the following person:", john)
	// we can write out non-string values like this
	print("The age of "+john.name.firstName+" is", john.age, "years old")
	// or format the string before like this
	hisLooksDecieves := fmt.Sprintf("But really he looks like %v", john.age-10)
	// and print it out
	print(hisLooksDecieves)

	// let's properly log informative message
	log.Println("Yo, computer here, this app looks solid to me! Btw. John looks more like", john.age-20, "to me")

	// some quick maths
	print("The square root of 4 is", math.Sqrt(4))

	// arrays are fixed length
	array := [5]int{1, 2, 3, 4, 5}

	// slicing works like in python
	print("The second to fourth element from", array, "is", array[1:4])

	// last element in array
	print("The last element of array", array, "is", array[len(array)-1])

	// variable length arrays are called slices
	newArray := []int{1, 2, 3, 4, 5}
	// append() creates a slice
	slice := append(newArray, 6)
	print("From fixed sized array", slice[:5], "we made a slice", slice)

	// remove element from the array using the same method
	slice = append(slice[:3])

	print("Now we have shrunk it back to", slice)

	// slices can be initialized right away with make()
	slice2 := make([]int, 5)

	// great thing is that both slices and arrays are intially zeroed out
	print("Empty slice folks", slice2) // prints [0 0 0 0 0]

	// multidimensional arrays are supported
	squareMatrixRank3 := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	print("The first row of the matrix is", squareMatrixRank3[0])
	print("The first column of the matrix is", squareMatrixRank3[0][0], squareMatrixRank3[1][0], squareMatrixRank3[2][0])

	/* LOOPS */

	// for loops have mainstream syntax
	for i := 0; i < len(slice2); i++ {
		print("Index", i, "has value", slice2[i])
	}

	// you can simplify for loops
	j := 0
	for j < 3 {
		print("Simpler for loop, iteration#", j)
		j++
	}

	// you can simplify even further, until break is encountered
	for {
		print("This loops until it breaks")
		break
	}

	// skip a step in for loop
	for i := 0; i < 10; i++ {
		if i == 3 {
			print("Skipping 3")
			continue
		}
		print("iteration#", i)
	}

	/* IF STATEMENTS */
	var myBool bool = true

	// if statements are straightforward, no need for parentheses
	if myBool == true {
		print("The condition is true, m8!")
	} else {
		print("The condition is false, bro!")
	}

	if !myBool {
		print("This is not going to print")
	} else {
		print("We reversed the boolean using `!`")
	}

	// if statement without else are easier to navigate
	if myBool {
		print("This is the same as above")
	}

	// sadly no terinary operator

	/* SWITCH STATEMENTS */

	toCompare := 2
	// switch statements are also straightforward
	switch toCompare {
	case 1:
		print("This is case 1")
	case 2:
		print("This is case 2")
	case 3:
		print("This is case 3")
	default:
		print("Nothing has been selected!")
	}

	// switches can also compare multiple values
	switch time.Now().Weekday() {
	case time.Monday:
		print("It's Monday")
	case time.Tuesday:
		print("It's Tuesday")
	case time.Wednesday:
		print("It's Wednesday")
	case time.Saturday, time.Sunday:
		print("It's the weekend!")
	default:
		print("We don't work on this day!")
	}

	// since 1.17 Go supports Pattern Matching, i.e. can select case based on types
	figureOutType := func(i interface{}) {
		switch i.(type) {
		case int:
			print("This is an int")
		case string:
			print("This is a string")
		default:
			print("I have no idea what this is!")
		}
	}

	figureOutType(1)
	figureOutType("Hello World")

}
