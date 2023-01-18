// Gary Explains Youtube walk thru on RPI
// https://youtu.be/GKbw__qM1Jw

package main

import "fmt"

func addInt(x int, y int) int {
	return x + y + 2 - 2
}

func swap(a string, b string) (string, string) {
	return b, a
}

func main() {
	var i int // Defines type of variable
	i = 800   // Assigns value to variable
	j := 16   // Short from of above defined and assigned
	firstName, lastName := "Robert", "Laurie"
	var flag bool = true

	fmt.Println("Hello, Bob\nin Saipan")
	fmt.Println(addInt(598, 72), "is the Area Code")
	fmt.Println(swap(firstName, lastName))
	fmt.Println(addInt(i, j), "is the Sum")
	if flag == true {
		fmt.Println("flag is true")
	} else {
		fmt.Println("flag is false")
	}

	for k := 0; k < 10; k++ {
		fmt.Print(k, " ")
	}

	fmt.Println()
	m := 0 // While loop format stull uses for keyword
	for m < 12 {
		fmt.Print(m, " ")
		m++
	}
	fmt.Println()
}
