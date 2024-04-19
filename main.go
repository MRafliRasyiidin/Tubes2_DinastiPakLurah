package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// func index(w http.ResponseWriter, r *http.Request) {
// fmt.Fprintln(w, "your mum")
// }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "")

		var t, err = template.ParseFiles(("template/main.html"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		t.Execute(w, map[string]string{"Text": "naw"})
	})

	// http.HandleFunc("/urmumgay", index)

	fmt.Println("starting web server at http://localhost:8080/")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(":8080", nil)
}
