package mocking

import "fmt"

func Countdown() {
	for count := 3; count >= 0; count-- {
		if count == 0 {
			fmt.Println("Go!")
			return
		}
		fmt.Printf("%d\n", count)
	}
}
