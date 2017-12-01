package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	day1()
}

func day1() {
	input, err := ioutil.ReadFile("day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	d1 := Day1_1(input)
	d2 := Day1_2(input)
	fmt.Printf("1.1 -> %d\n", d1)
	fmt.Printf("1.2 -> %d\n", d2)
}
