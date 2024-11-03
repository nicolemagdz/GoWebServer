package main

import (
	"embed"
	"net/http"
)

//go:embed syllabus.json
var jsonData embed.FS

func serveEmbeddedJSON(w http.ResponseWriter, r *http.Request) {
	data, err := jsonData.ReadFile("syllabus.json")
	if err != nil {
		http.Error(w, "Failed to read embedded JSON file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
