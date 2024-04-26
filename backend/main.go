package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Data struct {
    Start string `json:"startLink"`
    Target string `json:"TargetLink"`
}

func main() {
    //http.HandleFunc("/", searchHandler)
    http.HandleFunc("/search", searchHandler)
    log.Println("Server is running on http://localhost:3000")
    http.ListenAndServe(":3000", nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
    // Parse form data
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Error parsing form data", http.StatusInternalServerError)
        return
    }

    body, err := io.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Failed to read request body", http.StatusInternalServerError)
            return
        }

    // Get form values
    var data Data
        if err := json.Unmarshal(body, &data); err != nil {
            http.Error(w, "Failed to unmarshal request body", http.StatusBadRequest)
            return
        }

    startWord := data.Start
    targetWord := data.Target

    // Do something with the form data
    log.Printf("Received form data: Start Word = %s, Target Word = %s\n", startWord, targetWord)

    // You can send a response back to the client if needed
    fmt.Fprintf(w, "Received form data: Start Word = %s, Target Word = %s\n", startWord, targetWord)

    runWikiRace(startWord, targetWord, true);
}

func runWikiRace(start string, target string, searchType bool) {
    fmt.Println(start, target);
    algorithm(start, target)
}