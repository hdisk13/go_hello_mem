/*

On vacation this week.  I have some time to
check out 'go'.

I'm curious about what can be done with pointers
in this memory-safe and garbage-collected world.

*/

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

func singBottles(b int) {
	thisbeer := b
	nextbeer := b
	nextbeer -= 1

	if b > 1 {
		fmt.Println(strconv.Itoa(thisbeer) + " bottles of beer on the wall " + strconv.Itoa(thisbeer) + " bottles of beer")
		fmt.Println("take 1 down, pass it around, " + strconv.Itoa(nextbeer) + "bottles of beer on the wall!")
	} else {
		fmt.Println("1 bottle of beer on the wall, only 1 bottle of beer... Buurrrp!")
	}
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	a := 1

	fmt.Printf("Address is %p\n", &a)

	for i := 0; i < 10; i += 1 {
		myfunc2(i)
	}

	for b := 99; b > 0; b-- {
		singBottles(b)
	}

}
