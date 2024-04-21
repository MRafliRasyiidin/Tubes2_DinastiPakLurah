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
	"sync/atomic"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/thalesfsp/go-common-types/safeorderedmap"
)

func crawler(start string, target string) {
	queue := safeorderedmap.New[bool]()
	var currPath []string
	var continueSearch int32 = 1
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.Async(true),
	)
	done := make(chan bool)

	// c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 20})

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		if atomic.LoadInt32(&continueSearch) == 1 {
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
				!exists {
				queue.Add(h.Request.AbsoluteURL(link), true)
				fmt.Println(h.Request.AbsoluteURL(link), atomic.LoadInt32(&continueSearch))
			}
		}
		if link == "/wiki/"+target {
			atomic.StoreInt32(&continueSearch, 0)
			done <- true
			return
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
	})

	c.OnScraped(func(r *colly.Response) {
		if atomic.LoadInt32(&continueSearch) == 0 {
			return
		}
		fmt.Println("Finished", r.Request.URL)
		queue.Delete(r.Request.URL.String())
		var link string
		link = ""
		queue.Each(func(k string, e bool) {
			link = k
			currPath = append(currPath, link)
			if atomic.LoadInt32(&continueSearch) == 1 {
				c.Visit(link)
			}
		})
	})

	currPath = append(currPath, "https://en.wikipedia.org/wiki/"+start)
	c.Visit("https://en.wikipedia.org/wiki/" + start)
	go c.Wait()
	<-done
	// fmt.Println(currPath)
}

func main() {
	start := time.Now()
	crawler("Ring_of_Fire", "United_States")
	fmt.Println(time.Since(start))
}
