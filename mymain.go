package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("Huzzah, I ran on: %s\n", now.Format(time.RFC3339))

	// reference item in "internal" package
	//fmt.Printf("I can see: %s\n", privatestuff.MyPrivateConst)
}

func myUnusedFunction() {
	// nobody home
}

func secondUnused() {
	// also nothing happening here
}
