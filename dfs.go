package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func crawler(start string, target string) {
	maxDepth := 1
	for {
		if dfs(start, target, maxDepth, 1, []string{}) {
			break
		}
		maxDepth++
	}
}

var continueSearch bool = true

func dfs(start string, target string, maxDepth, depth int, currPath []string) bool {
	if depth > maxDepth || !continueSearch {
		return false
	}

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"))

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		if continueSearch && !isInPath(link, currPath) {
			if link == "/wiki/"+strings.ReplaceAll(target, " ", "_") {
				continueSearch = false
				fmt.Println("Found in depth:", depth)
			}
			if len(currPath) < maxDepth && !strings.HasPrefix(link, "#") && !strings.HasPrefix(link, "http") {
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
					(link != "/wiki/Main_Page") {
					fmt.Printf("%s - depth: %d\n", h.Request.AbsoluteURL(link), depth)
					dfs(extractTitle(link), target, maxDepth, depth+1, append(currPath, link))
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

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	if continueSearch {
		c.Visit("https://en.wikipedia.org/wiki/" + strings.ReplaceAll(start, " ", "_"))
		c.Wait()
	}
	return !continueSearch
}

func isInPath(link string, path []string) bool {
	for _, p := range path {
		if link == p {
			return true
		}
	}
	return false
}

func extractTitle(url string) string {
	parts := strings.Split(url, "/")
	title := parts[len(parts)-1]
	title = strings.ReplaceAll(title, "_", " ")
	return title
}
