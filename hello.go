package main

import (
	"fmt"
	"strconv"

	"rsc.io/quote"
)

func mem_print(addy *func()) {
	fmt.Println("Treading on thin ice!")
	fmt.Printf("Address is %p\n", &addy)
}

func myfunc1() {
	fmt.Println("Hello f1")
	a := 1
	b := 1
	c := 1

	fmt.Printf("Address is %p\n", &a)
	fmt.Printf("Address is %p\n", &b)
	fmt.Printf("Address is %p\n", &c)

}

func myfunc2(loopcount int) {
	fmt.Println("Hello f2 : loop " + strconv.Itoa(loopcount))

	a := 1
	b := 1
	c := 1

	fmt.Printf("Address is %p\n", &a)
	fmt.Printf("Address is %p\n", &b)
	fmt.Printf("Address is %p\n", &c)
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	a := 1
	b := 1
	c := 1

	fmt.Printf("Address is %p\n", &a)
	fmt.Printf("Address is %p\n", &b)
	fmt.Printf("Address is %p\n", &c)
	for i := 0; i < 10; i += 1 {
		myfunc2(i)
	}

}
