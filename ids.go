package main
//
//	import (
//		"fmt"
//		"strings"
//
//		"github.com/gocolly/colly/v2"
//	)
//
//var continueSearch bool = true
//var count int = 1
//
//func ids(start string, target string, maxDepth, depth int) {
//	if depth > maxDepth || !continueSearch {
//		return
//	}
//
//	c := colly.NewCollector(
//		colly.AllowedDomains("en.wikipedia.org"))
//
//	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
//		link := h.Attr("href")
//		if continueSearch && !isInPath(link, p) {
//			if link == "/wiki/"+strings.ReplaceAll(target, " ", "_") {
//				continueSearch = false
//				fmt.Println("Found in depth:", depth)
//			}
//			if !strings.HasPrefix(link, "#") && !strings.HasPrefix(link, "http") {
//				if strings.HasPrefix(link, "/wiki/") &&
//					!strings.Contains(link, "File:") &&
//					!strings.Contains(link, "Help:") &&
//					!strings.Contains(link, "Category:") &&
//					!strings.Contains(link, "Wikipedia:") &&
//					!strings.Contains(link, "Talk:") &&
//					!strings.Contains(link, "Special:") &&
//					!strings.Contains(link, "Portal:") &&
//					!strings.Contains(link, "Template:") &&
//					!strings.Contains(link, "MediaWiki:") &&
//					!strings.Contains(link, "User:") &&
//					!strings.Contains(link, "_talk:") &&
//					(link != "/wiki/Main_Page") && extractTitle(link) != start && !isInPath("https://en.wikipedia.org" + link, visited){
//					p = append(p, "https://en.wikipedia.org" + link)
//					fmt.Printf("%s - depth: %d\n", "https://en.wikipedia.org" + link, depth)
//					if depth != 1 {
//						depth -= 1
//						visited = append(visited, "https://en.wikipedia.org" + link)
//						c.Visit("https://en.wikipedia.org" + link)
//					}
//				}
//			}
//		}
//	})
//
//	c.OnRequest(func(r *colly.Request) {
//		fmt.Println("Visiting", r.URL)
//	})
//
//	c.OnError(func(_ *colly.Response, err error) {
//		fmt.Println("Something went wrong:", err)
//	})
//
//	c.OnResponse(func(r *colly.Response) {
//		// fmt.Println("Visited", r.Request.URL)
//	})
//
//	c.OnScraped(func(r *colly.Response) {	// run ketika seluruh link di web udah disimpen di objek colly
//		fmt.Println("Finished", r.Request.URL)
//		count += 1
//		fmt.Println("ini count:",count)
//		if depth == 1 {
//			depth = maxDepth
//		}
//	})
//
//	c.Visit("https://en.wikipedia.org/wiki/" + strings.ReplaceAll(start, " ", "_"))
//	count += 1
//	c.Wait()
//}
//var p []string
//var visited []string
//
//func main() {
//	for i := 1; i < 10; i++ {
//		ids("Ring of Fire", "Car", i, i) 
//		p = p[:0]
//		visited = visited[:0]
//		//fmt.Println("ini i", i)
//	}
//	for i := 0; i < len(p); i++ {
//		fmt.Println(p[i])
//		fmt.Println("v",visited[i])
//	}
//	fmt.Println(continueSearch)
//}
//