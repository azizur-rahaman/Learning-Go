package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	func booksHandler(w http.http.ResponseWriter, r *http.Request){
		switch r.Method {
			case "GET":
				fmt.Fprintf(w, "List of books")
			case "POST":
				fmt.Fprintf(w, "Create a new book")
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}

    func sortedBooks(w http.http.ResponseWriter, r *http.Request){
		switch r.Method{
			case "GET":
				genre := r.URL.Query().Get("genre")
				if genre != "" {
					fmt.Fprintf(w, "Books in genre: %s", genre)
				}
				body, err := io.ReadAll(r.Body)
				if err != nil {
					http.Error(w, "Failed to read body", http.StatusInternalServerError)
					return
				}
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
	
}
