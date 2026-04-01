package main

import "fmt"

// This program demonstrates the use of keywords in Go.
// use break        default      func         interface    select
// case         defer        go           map          struct
// chan         else         goto         package      switch
// const        fallthrough  if           range        type
// continue     for          import       return       var
func main() {
	// This is a comment
	var x int = 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is not greater than 5")
	}
	cats := []string{"Whiskers", "Fluffy", "Mittens"}
	for i, cat := range cats {
		fmt.Println(i, cat)
	}
	go func() {
		fmt.Println("Hello from a goroutine")
	}()
	switch x {
	case 1:
		fmt.Println("x is 1")
	case 10:
		fmt.Println("x is 10")
	default:
		fmt.Println("x is something else")
	}
}
