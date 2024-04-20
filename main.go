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

	"github.com/gocolly/colly/v2"
)

func crawler() {
	queue := make(map[string]bool)
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.Async(true))
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})
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
			("/wiki/Main_Page" != link) &&
			!queue["https://en.wikipedia.org"+link] {
			queue["https://en.wikipedia.org"+link] = true
			fmt.Println("https://en.wikipedia.org" + link)
		}
		h.Request.Visit(link)
	})
	queue["https://en.wikipedia.org/wiki/Gibran_Rakabuming"] = true
	c.Visit("https://en.wikipedia.org/wiki/Gibran_Rakabuming")
	c.Wait()
}

func main() {
	crawler()
}
