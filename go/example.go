package main

import (
	// input/output
	"fmt"
	// error messages / handling
	"errors"
	// math.. is math
	"math"

)

// create own types (struct) outside of functions
type person struct {
	name string
	age int
}

func main() {
	
	fmt.Println("\n Welcome to little Golang")
	// variables

	fmt.Println("\n First, print some integers and some text")
	x := 5
	y := 7
	var z int = 10
	sum := x+y
	fmt.Println(sum)
	fmt.Println(z)
	fmt.Println("Hello Go")
	
	// Conditions

	fmt.Println("\n Next, check if x is this or that")

	if x > 6 {
		fmt.Println("x is greater 6")
	} else if x < 2 {
		fmt.Println("x is smaller 2")
	} else {
		fmt.Println("x is between 2 and 6")
	}

	// array

	fmt.Println("\n Now, let's look at arrays")

	var a [5]int
	a[2] = 7
	fmt.Println(a)

	b := [5]int{5,4,3,2,1}
	fmt.Println(b)

	fmt.Println("\n Ok.. this two arrays where fixed in there length. Now get to variable length")
	c := []int{9,8,7,6,5}
	fmt.Println(c)
	c = append(c,13)
	fmt.Println(c)

	// maps

	fmt.Println("\n Maps are like dictionary's in python..")

	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])

	delete(vertices, "square")
	fmt.Println(vertices)
	
	// loops - 	only one loop (for)

	fmt.Println("\n Let's do some loopings")
	fmt.Println("\n Start with for loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	fmt.Println("\n Do while..")
	j := 0
	for j<5 {
		fmt.Println(j)
		j++
	}

	// loop over elements in an array
	fmt.Println("\n Elements in array")
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	// it also works with a map
	fmt.Println("\n Elements in map")
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}

	// now lets use other functions
	fmt.Println("\n Call functions to do what they should do")
	result := sumfunc(2,3)
	fmt.Println(result)

	fmt.Println("\n Now with more return values")
	result2, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}

	result3, err := sqrt(-16)
	
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result3)
		}

	// use self-created types / struct

	fmt.Println("\n Working with self created types / struct")
	p := person{name: "Jake", age: 23}
	fmt.Println(p)
	fmt.Println(p.age)

	// pointers
	fmt.Println("\n Point to the pointers")
	k := 7
	inc(&k)
	fmt.Println(k)

	}

func inc (x *int) {
	*x++
}

// x,y parameters, define type of return-value
func sumfunc(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		// return only error
		return 0, errors.New("Undefined for negative numbers")
	}

	// all is fine, return no error
	return math.Sqrt(x), nil
}
