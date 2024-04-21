package main

import "fmt"

func maina() {
	for i := 0; i < 9 && continueSearch; i++ {
		crawler("Ring of Fire", "Street dance", i)
		fmt.Println(continueSearch)
	}
}
