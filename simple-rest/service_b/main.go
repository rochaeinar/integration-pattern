package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Document struct {
	Id   string
	Name string
	Size int
}

func main() {
	fmt.Println("Starting REST API Service-B")
	handleRequest()
}

func handleRequest() {
	http.HandleFunc("/documents", getDocuments)
	log.Fatal(http.ListenAndServe(":9100", nil))
}
func getDocuments(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Welcome to my RESST service")
	var docs []Document
	docs = append(docs, Document{Id: "Doc 1", Name: "Report-quarter-1.pdf", Size: 1899})
	docs = append(docs, Document{Id: "Doc 2", Name: "Report-quarter-2.pdf", Size: 5419})
	docs = append(docs, Document{Id: "Doc 3", Name: "Report-quarter-3.pdf", Size: 954})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docs)
}
