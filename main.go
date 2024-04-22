// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"net/http"
// )

// // func index(w http.ResponseWriter, r *http.Request) {
// // fmt.Fprintln(w, "your mum")
// // }

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "")

// 		var t, err = template.ParseFiles(("template/main.html"))
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return
// 		}
// 		t.Execute(w, map[string]string{"Text": "naw"})
// 	})

// 	// http.HandleFunc("/urmumgay", index)

// 	fmt.Println("starting web server at http://localhost:8080/")
// 	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
// 	http.ListenAndServe(":8080", nil)
// }

package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

func crawler(start string, target string) {
	// queue := safeorderedmap.New[bool]()
	var queue sync.Map
	var path sync.Map
	// path := safeorderedmap.New[string]()
	// path := orderedmap.NewOrderedMap[string, string]()
	done := make(chan bool)
	// urlChan := make(chan string)
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.Async(true),
		colly.CacheDir("./cache"),
	)
	c.AllowURLRevisit = false

	// Wtf is even Parallelism: 1000?? Me brainrot big number equals good
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 100, Delay: 20 * time.Millisecond})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		// I need to track how many links are fully visited
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		visited, _ := c.HasVisited(link)
		_, exists := queue.Load(h.Request.AbsoluteURL(link))
		_, exists2 := path.Load(h.Request.AbsoluteURL(link))
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
			link != "/wiki/Main_Page" &&
			h.Request.AbsoluteURL(link) != h.Request.URL.String() &&
			!exists && !visited {
			queue.Store(h.Request.AbsoluteURL(link), true)
			// path.Add(h.Request.AbsoluteURL(link), h.Request.URL.String())
			if !exists2 {
				path.Store(h.Request.AbsoluteURL(link), h.Request.URL.String())
			}
			// fmt.Println(h.Request.AbsoluteURL(link))
		}
		if link == "/wiki/"+target {
			path.Store(h.Request.AbsoluteURL(link), h.Request.URL.String())
			// path.Add(h.Request.AbsoluteURL(link), h.Request.URL.String())
			done <- true
			return
		}
	})

	// Synced
	// go func() {
	// 	for url := range urlChan {
	// 		c.Visit(url)
	// 	}
	// }()

	c.OnScraped(func(r *colly.Response) {
		// fmt.Println("Finished", r.Request.URL)
		queue.Delete(r.Request.URL.String())
		var link string
		link = ""
		queue.Range(func(key, value any) bool {
			link = key.(string)
			c.Visit(link)
			return true
		})
		// queue.Each(func(k string, e bool) {
		// link = k
		// c.Visit(link)
		// urlChan <- k
		// })
	})

	c.Visit("https://en.wikipedia.org/wiki/" + start)
	go c.Wait()
	<-done

	// close(urlChan)

	key := "https://en.wikipedia.org/wiki/" + target
	for key != "https://en.wikipedia.org/wiki/"+start {
		fmt.Println("Key : ", key)
		value, _ := path.Load(key)
		key = value.(string)
	}
	// fmt.Println("Key : ", key)
	// path.Each(func(key, value string) {
	// 	fmt.Println("PATH:", key, value)
	// })
}

func pageloader(start string) {
	done := make(chan bool)
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.Async(false),
		colly.CacheDir("./cache"),
	)
	c.AllowURLRevisit = false

	// Wtf is even Parallelism: 1000?? Me brainrot big number equals good
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 1, Delay: 100 * time.Millisecond})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		// I need to track how many links are fully visited
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
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

			h.Request.AbsoluteURL(link) != h.Request.URL.String() &&
			link != "/wiki/Main_Page" {
			fmt.Println(link)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		done <- true
	})

	c.Visit("https://en.wikipedia.org/wiki/" + start)
	go c.Wait()
	<-done
}

func main() {
	start := time.Now()
	// pageloader("Ring_of_Fire")
	crawler("MiHoYo", "Intel_8080")
	fmt.Println(time.Since(start))
}
