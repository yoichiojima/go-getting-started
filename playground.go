package main

import "fmt"

// variables
func testVariable() {
	var a int = 1
	var b int = 100

	fmt.Println("variables")

	fmt.Println(a, b)
}

// calculation
func testCalculation() {
	var a int = 100
	var b int = 200

	fmt.Println("calculation")

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
}

// array
func testArray() {
	var a [3]int = [3]int{1, 2, 3}
	var b [3]int = [3]int{4, 5, 6}

	fmt.Println("array")

	fmt.Println(a)
	fmt.Println(b)
}

// map
func testMap() {
	m := map[string]int{"apple": 100, "banana": 200}

	fmt.Println("map")
	fmt.Println(m)

	fmt.Println(m["apple"])
}

func testIf() {
	const a int = -1

	if a > 0 {
		fmt.Println("a is larger than 0")
	} else if a < 0 {
		fmt.Println("a is smaller than 0")
	} else {
		fmt.Println("a is 0")
	}
}

func main() {
	// testVariable()
	// testCalculation()
	// testArray()
	// testMap()
	testIf()
}
