package main

import (
	"encoding/json"
	//"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Data struct {
    Start string `json:"startLink"`
    Target string `json:"TargetLink"`
    Method string `json:"searchType"`
}


func main() {
    //http.HandleFunc("/", index)
    http.HandleFunc("/search", searchHandler)
    log.Println("Server is running on http://localhost:3000")
    http.ListenAndServe(":3000", nil)
}

//func index(w http.ResponseWriter, r *http.Request) {
//    http.ServeFile(w, r, "src/index.html")
//}

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
    searchMethod := data.Method
    startWord = strings.ReplaceAll(startWord, " ", "_")
    targetWord = strings.ReplaceAll(targetWord, " ", "_")
    // Do something with the form data
    log.Printf("Received form data: Start Word = %s, Target Word = %s\n", startWord, targetWord)

    // You can send a response back to the client if needed
    //fmt.Fprintf(w, "Received form data: Start Word = %s, Target Word = %s\n", startWord, targetWord)

    //runWikiRace(startWord, targetWord, searchMethod);
    bfsMethod := true
    if searchMethod == "IDS" {
        bfsMethod = false
    }
    path_result := runAlgorithm(startWord, targetWord, bfsMethod)
    // Set the content type header to JSON
    w.Header().Set("Content-Type", "application/json")

    // Write the JSON-encoded result to the response body
    w.Write(path_result)
}

//func runWikiRace(start string, target string, searchType string) {
//    bfsMethod := true
//    if searchType == "IDS" {
//        bfsMethod = false
//    }
//    path_result := runAlgorithm(start, target, bfsMethod)
//    // Set the content type header to JSON
//    w.Header().Set("Content-Type", "application/json")
//
//    // Write the JSON-encoded result to the response body
//    w.Write(result)
//    //fmt.Println(searchType)
//    //fmt.Println(start, target)
//}