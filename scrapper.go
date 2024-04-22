package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

var continueSearch bool = true

func scraper(visitLink string, target string, depth int) []string {

	var link_result []string

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"))

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		if continueSearch {
			if link == "/wiki/"+strings.ReplaceAll(target, " ", "_") {
				continueSearch = false
			}
			if !strings.HasPrefix(link, "#") && !strings.HasPrefix(link, "http") {
				var a string = "https://en.wikipedia.org" + link
				if strings.HasPrefix(link, "/wiki/") &&
					!strings.Contains(link, "File:") &&
					!strings.Contains(link, "Help:") &&
					!strings.Contains(link, "Category:") &&
					!strings.Contains(link, "Wikipedia:") &&
					!strings.Contains(link, "Talk:") &&
					!strings.Contains(link, "Special:") &&
					!strings.Contains(link, "Portal:") &&
					!strings.Contains(link, "Template:") &&
					!strings.Contains(link, "MediaWiki:") &&
					!strings.Contains(link, "User:") &&
					!strings.Contains(link, "_talk:") &&
					(link != "/wiki/Main_Page") && !isIn(link_result, a){
						link_result = append(link_result, "https://en.wikipedia.org" + link)
						fmt.Println("ini url", "https://en.wikipedia.org" + link)
				}
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		// fmt.Println("Visited", r.Request.URL)
	})

	c.OnScraped(func(r *colly.Response) {	// run ketika seluruh link di web udah disimpen di objek colly
		fmt.Println("Finish Scraped", r.Request.URL)
		for _,url := range link_result {
			if continueSearch && depth != 1 {
				scraper(url, target, depth - 1)
			}
			if !continueSearch {
				break
			}
		}

	})

	if continueSearch {
		c.Visit(visitLink)
		c.Wait()
	}
	return link_result
}

func isIn(list []string, check string) bool {
	for _,el := range list {
		if el == check {
			return true
		}
	}
	return false
}