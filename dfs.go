package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func crawlerOLD(start string, target string) {
	maxDepth := 1
	for {
		if dfsOld(start, target, maxDepth, 1, []string{}) {
			break
		}
		maxDepth++
	}
}

var continueSearch bool = true

func dfsOld(start string, target string, maxDepth, depth int, currPath []string) bool {
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
					newPath := append([]string(nil), currPath...)
					dfsOld(extractTitle(link), target, maxDepth, depth+1, append(newPath, link))
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

func crawler(start, target string) ([]string, bool) {
	maxDepth := 3
	path, found := ids(start, target, maxDepth)
	if found {
		return path, true
	}
	return nil, false
}

func ids(start, target string, maxDepth int) ([]string, bool) {
	for depth := 1; depth <= maxDepth; depth++ {
		fmt.Printf("Searching at depth %d...\n", depth)
		visited := make(map[string]bool)
		path, found := dfs(start, target, depth, visited, []string{start}, false)
		if found {
			return path, true
		}
	}
	return nil, false
}

func dfs(start, target string, depth int, visited map[string]bool, path []string, isFound bool) ([]string, bool) {
	if isFound {
		return path, true
	}
	if depth == 0 {
		return nil, false
	}
	if start == target {
		return path, true
	}
	if visited[start] {
		return nil, false
	}

	visited[start] = true

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"))

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if !isFound {
			if link == "/wiki/"+strings.ReplaceAll(target, " ", "_") {
				isFound = true
				fmt.Println("Found in depth: ", depth)
			}
			if !strings.HasPrefix(link, "#") && !strings.HasPrefix(link, "http") {
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
					newLink := extractTitle(link)
					newPath := append(path, newLink)
					var result []string
					result, isFound = dfs(newLink, target, depth-1, visited, newPath, isFound)
					if isFound {
						fmt.Println("Path: found")
						fmt.Print("[")
						for _, link := range result {
							fmt.Print(link + "->")
						}
						fmt.Print("]")
						fmt.Println()
						return
					}
				}
			}
		}
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %s\n", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		// fmt.Println("Visited", r.Request.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		// fmt.Println("Finished", r.Request.URL)
	})

	if !isFound {
		c.Visit("https://en.wikipedia.org/wiki/" + strings.ReplaceAll(start, " ", "_"))
		c.Wait()
	}

	return nil, false
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
