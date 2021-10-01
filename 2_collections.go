package main

import "fmt"

// https://pkg.go.dev/container
// https://github.com/emirpasic/gods/

func printArray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func printSlices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
	//---------------------------------------
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[:1]
	c := names[1:3]
	fmt.Println(a, b, c)
	//---------------------------------------
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
}

func printSlicesLens() {
	printSlice := func(s []int) {
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSliceMake() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)
	//--------------------------
	var s []int
	printSlice2(s)
	// append works on nil slices.
	s = append(s, 0)
	printSlice2(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice2(s)
}

func printLoops() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func printMap() {
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{}
	fmt.Println(m["Bell Labs"])

	n := map[string]int{
		"Answer": 1,
	}

	fmt.Println(n["Answer"])

	//----------------------------------
	n["Answer"] = 42
	fmt.Println("The value:", n["Answer"])

	n["Answer"] = 48
	fmt.Println("The value:", n["Answer"])

	delete(n, "Answer")
	fmt.Println("The value:", n["Answer"])

	v, ok := n["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// ----------------------------
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
