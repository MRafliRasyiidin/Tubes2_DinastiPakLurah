package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Page struct {
	Title string
	URL   string
	Links []string
}

func breadthFirstScrapper(url string, word string) {
	c := colly.NewCollector()

	var (
		isFound bool
		queue   []string
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		if !strings.HasPrefix(link, "http") {
			link = "https://id.wikipedia.org" + link
		}

		if !strings.Contains(link, "wikipedia.org") || !strings.Contains(link, "/wiki/") {
			return
		}

		if strings.Contains(link, "Templat") {
			return
		}

		if strings.Contains(link, "Kategori:") {
			return
		}

		fmt.Println("Link found:", link)
		if strings.Contains(link, word) {
			fmt.Println("Found the specified word at:", link)
			isFound = true
			return
		}

		queue = append(queue, link)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	queue = append(queue, url)

	for len(queue) > 0 && !isFound {
		url := queue[0]
		queue = queue[1:]

		fmt.Println("Visiting:", url)
		c.Visit(url)
	}
	if isFound {
		fmt.Println("KEtEMU")
	}
}
