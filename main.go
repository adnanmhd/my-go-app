package main

import "fmt"

func main2() {
	fmt.Println("Hello World")
	fmt.Println("Pyramid 1: ")
	getPyramid()
	fmt.Println("Pyramid 2: ")
	getPyramid2()
	fmt.Println("Pyramid 3: ")
	getPyramid3()
	fmt.Println("Pyramid 4: ")
	getPyramid4()
	var length int = 50
	fmt.Println("Fibonacci, length:", length)
	fiboancci(length)
}

func getPyramid() {
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func getPyramid2() {
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func getPyramid3() {
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func getPyramid4() {
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(" ")
		}
		for k := i; k < 10; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func fiboancci(limit int) {

	var bil1 int64 = 0
	var bil2 int64 = 1
	for i := 0; i < limit; i++ {
		var countNumber int64
		fmt.Print(bil1, ", ")
		countNumber = bil1 + bil2
		bil1 = bil2
		bil2 = countNumber
	}
}
