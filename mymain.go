package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("Huzzah, I ran on: %s\n", now.Format(time.RFC3339))

	//myUnusedFunction()
}

func myUnusedFunction() {
	// nobody home
}
