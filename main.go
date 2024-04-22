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

	"github.com/elliotchance/orderedmap/v2"
	"github.com/gocolly/colly/v2"
)

func crawler(start string, target string) {
	var mutex sync.Mutex
	queue := orderedmap.NewOrderedMap[string, any]()
	path := orderedmap.NewOrderedMap[string, any]()
	queueChild := orderedmap.NewOrderedMap[string, any]()
	found := false
	done := make(chan bool)

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.Async(true),
		// colly.CacheDir("./cache"),
	)
	c.AllowURLRevisit = false

	// Wtf is even Parallelism: 1000?? Me brainrot big number equals good
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 100, RandomDelay: 25 * time.Millisecond})

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
		// time.Sleep(1 * time.Millisecond)
		link := h.Attr("href")
		visited, _ := c.HasVisited(link)
		mutex.Lock()
		_, exists := queue.Get(h.Request.AbsoluteURL(link))
		mutex.Unlock()
		mutex.Lock()
		_, exists2 := path.Get(h.Request.AbsoluteURL(link))
		mutex.Unlock()
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
			mutex.Lock()
			queueChild.Set(h.Request.AbsoluteURL(link), true)
			mutex.Unlock()
			// fmt.Println(h.Request.AbsoluteURL(link))
			if !exists2 {
				mutex.Lock()
				path.Set(h.Request.AbsoluteURL(link), h.Request.URL.String())
				mutex.Unlock()
			}
		}
		if link == "/wiki/"+target {
			mutex.Lock()
			path.Set(h.Request.AbsoluteURL(link), h.Request.URL.String())
			mutex.Unlock()
			found = true
			done <- true
			return
		}
	})

	c.OnScraped(func(r *colly.Response) {
		// fmt.Println("Finished", r.Request.URL)
		mutex.Lock()
		queue.Delete(r.Request.URL.String())
		queue, queueChild = queueChild, queue
		queueChild = orderedmap.NewOrderedMap[string, any]()
		mutex.Unlock()
	})

	queue.Set("https://en.wikipedia.org/wiki/"+start, true)

queueIteration:
	for {
		mutex.Lock()
		for el := queue.Front(); el != nil; el = el.Next() {
			c.Visit(el.Key)
		}
		mutex.Unlock()
		if found {
			break queueIteration
		}

	}
	go c.Wait()
	<-done

	key := "https://en.wikipedia.org/wiki/" + target
	for key != "https://en.wikipedia.org/wiki/"+start {
		fmt.Println("Key : ", key)
		value, _ := path.Get(key)
		key = value.(string)
	}
	// fmt.Println("Key : ", key)
	// path.Each(func(key, value string) {
	// 	fmt.Println("PATH:", key, value)
	// })
}

func main() {
	start := time.Now()
	// pageloader("Ring_of_Fire")
	crawler("MiHoYo", "Intel_8080")
	fmt.Println(time.Since(start))
}
