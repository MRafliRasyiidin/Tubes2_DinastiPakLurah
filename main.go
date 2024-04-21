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
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/thalesfsp/go-common-types/safeorderedmap"
)

func crawler(start string, target string) {
	queue := safeorderedmap.New[bool]()
	path := safeorderedmap.New[string]()
	done := make(chan bool)
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.Async(true),
	)
	c.AllowURLRevisit = false

	// Wtf is even Parallelism: 1000?? Me brainrot big number equals good
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 1000, Delay: 20 * time.Millisecond})

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
			link != "/wiki/Main_Page" &&
			!exists && !visited {
			queue.Add(h.Request.AbsoluteURL(link), true)
			path.Add(h.Request.AbsoluteURL(link), h.Request.URL.String())
			// fmt.Println(h.Request.AbsoluteURL(link))
		}
		if link == "/wiki/"+target {
			path.Add(h.Request.AbsoluteURL(link), h.Request.URL.String())
			done <- true
			return
		}
	})

	c.OnScraped(func(r *colly.Response) {
		// fmt.Println("Finished", r.Request.URL)
		queue.Delete(r.Request.URL.String())
		var link string
		link = ""
		queue.Each(func(k string, e bool) {
			link = k
			c.Visit(link)
		})
	})

	c.Visit("https://en.wikipedia.org/wiki/" + start)
	go c.Wait()
	<-done

	key := "https://en.wikipedia.org/wiki/" + target
	for key != "https://en.wikipedia.org/wiki/"+start {
		fmt.Println("Key : ", key)
		value, _ := path.Get(key)
		key = value
	}
	fmt.Println("Key : ", key)
	// path.Each(func(key, value string) {
	// 	fmt.Println("PATH:", key, value)
	// })
}

func main() {
	start := time.Now()
	crawler("MiHoYo", "Intel_8080")
	fmt.Println(time.Since(start))
}
