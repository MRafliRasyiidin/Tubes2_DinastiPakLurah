package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func Scrapper() {
	c := colly.NewCollector()
	var visitCount int = 0
	var links []string
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		visitCount++
		if visitCount > 5 {
			return
		}
		link := e.Attr("href")
		fmt.Println("Link found:", link)

		links = append(links, link)
		e.Request.Visit(link)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://id.wikipedia.org/wiki/Halaman_Utama")
}

func depthScrapper(url string, depth int, word string) {
	if depth > 5 {
		return
	}

	c := colly.NewCollector()

	var visitCount int = 0

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if depth >= 5 {
			return
		}

		link := e.Attr("href")
		fmt.Println("Link found:", link)

		if strings.Contains(link, word) {
			fmt.Println("Found the specified word at depth:", depth)
			return
		}

		if depth < 5 {
			depthScrapper(url, depth+1, word)
		}
		e.Request.Visit(link)

		visitCount++
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(url)
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
