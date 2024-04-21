package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
	orderedmap "github.com/wk8/go-ordered-map"
)

func crawler(start string, target string, depth int) {
	queue := orderedmap.New()
	var currPath []string
	continueSearch := true
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"))
	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		if link == "/wiki/"+target {
			continueSearch = false
		}
		_, exists := queue.Get(h.Request.AbsoluteURL(link))
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
			(link != "/wiki/Main_Page") &&
			!exists && len(currPath) <= depth {
			queue.Set(h.Request.AbsoluteURL(link), true)
			fmt.Println(h.Request.AbsoluteURL(link))
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

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		queue.Delete(r.Request.URL.String())
		var link string
		link = ""
		for key := queue.Oldest(); key != nil; key = key.Next() {
			link = key.Key.(string)
			break
		}
		currPath = currPath[:len(currPath)-1]
		if continueSearch && len(currPath) <= depth {
			currPath = append(currPath, link)
			c.Visit(link)
		}
	})

	currPath = append(currPath, "https://en.wikipedia.org/wiki/"+strings.ReplaceAll(start, " ", "_"))
	c.Visit("https://en.wikipedia.org/wiki/" + strings.ReplaceAll(start, " ", "_"))
	c.Wait()
	fmt.Println(currPath)
}
