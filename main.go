package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]string{"message": fmt.Sprintf("Received data with title: %s", data["title"])}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
