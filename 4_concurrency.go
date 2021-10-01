package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// v := Vertex{3, 4}
	// v.Abs()
	// fmt.Printf("%v", v.X)
	//fmt.Println()
	// v.Abs2()
	// fmt.Printf("%v", v.X)

	// ------------------
	// var a Abser
	// v := Vertex{3, 4}
	// // a = v
	// // a = &v
	// fmt.Print(a)

	// ------------------

	//printArray()

	// ------------------
	// printSlices()

	// -------------------
	//printSlicesLens()

	// ------------------

	// go count("sheep")
	// count("fish")

	//--------------------

	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {
	// 	count("sheep")
	// 	wg.Done()
	// }()

	// wg.Wait()

	// ---------------------

	// c := make(chan string, 50)
	// go chanCount("sheep", c)

	// msg := <-c
	// fmt.Println(msg)
	// // for msg := range c {
	// // 	fmt.Println(msg)
	// // }

	// -------------------------------

	// c1 := make(chan string)
	// c2 := make(chan string)

	// go func() {
	// 	for {
	// 		time.Sleep(time.Millisecond * 500)
	// 		c1 <- "Every 500ms"
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		time.Sleep(time.Second * 2)
	// 		c2 <- "Every two seconds"
	// 	}
	// }()

	// for {
	// 	select {
	// 	case msg1 := <-c1:
	// 		fmt.Println(msg1)
	// 	case msg2 := <-c2:
	// 		fmt.Println(msg2)
	// 	}
	// }

	// ------------------------------
	// ShowDefers()

	// ------------------------------

	// f := createFile("./defer.txt")
	// defer closeFile(f)
	// writeFile(f)
}

func count(thing string) {
	for i := 1; i < 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func chanCount(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	//close(c) // play with this to see what happends
}

func ShowDefers() { //LIFO
	defer Hello(1)
	defer Hello(2)
	defer Hello(3)
}
func Hello(i int) {
	fmt.Printf("%v", i)
	fmt.Println()
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
