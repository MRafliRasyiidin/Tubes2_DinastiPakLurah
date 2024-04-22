package main

import (
	"fmt"
	"time"
)

const MAX_DEPTH int = 4

func main() {
	start := time.Now()
	var first_depth []string = scraper("https://en.wikipedia.org/wiki/Ring_of_Fire", "Nigger", 1)
	for depth := 2; depth < MAX_DEPTH; depth++ {
		for _,i := range first_depth {
			if continueSearch {
				var res []string = scraper(i, "Nigger", depth)
				_ = res
			} else {
				break
			}
		}
	}
	
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(continueSearch)
}