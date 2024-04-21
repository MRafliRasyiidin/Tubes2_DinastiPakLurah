package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
	orderedmap "github.com/wk8/go-ordered-map"
)

func crawler(start string, target string, maxDepth int) {
	dfs(start, target, maxDepth, 1, []string{})
}

func dfs(start string, target string, maxDepth, depth int, currPath []string) bool {
	if depth > maxDepth {
		return false
	}

	queue := orderedmap.New()
	continueSearch := true

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"))

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		if link == "/wiki/"+strings.ReplaceAll(target, " ", "_") {
			continueSearch = false
			ansDepth := len(currPath)
			fmt.Println("Found in depth:", ansDepth)
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
			!exists && len(currPath) < maxDepth {
			queue.Set(h.Request.AbsoluteURL(link), true)
			fmt.Printf("%s - depth: %d\n", h.Request.AbsoluteURL(link), depth)
			dfs(start, target, maxDepth, depth+1, append(currPath, h.Request.AbsoluteURL(link)))
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
	})

	c.Visit("https://en.wikipedia.org/wiki/" + strings.ReplaceAll(start, " ", "_"))
	c.Wait()
	return !continueSearch
}
