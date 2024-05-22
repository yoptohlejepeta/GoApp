package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "go Go GO!"})
	})
	http.HandleFunc("/query", queryHandler)
	http.HandleFunc("/path/", pathHandler)
	http.HandleFunc("/body", bodyHandler)

	log.Println("SERVER RUNNING | PORT: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")

	response := map[string]string{"message": fmt.Sprintf("Hello, %s!", name)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/path/"):] // Extract the ID from the URL path

	response := map[string]string{"message": fmt.Sprintf("Requested resource ID: %s", id)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func bodyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var article Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": fmt.Sprintf("Received article with title: %s", article.Title),
		"article": map[string]string{
			"title":   article.Title,
			"content": article.Content,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
