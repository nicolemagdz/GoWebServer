package main

import (
	"embed"
	"fmt"
	"net/http"
)

//go:embed syllabus.json
var syllabusData embed.FS

func syllabusHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		serveSyllabus(w)
	case "DELETE":
		deleteSyllabus(w)
	case "POST":
		createSyllabus(w)
	case "PUT":
		updateSyllabus(w)	
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func serveSyllabus(w http.ResponseWriter) {
	data, err := syllabusData.ReadFile("syllabus.json")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read syllabus data: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func deleteSyllabus(w http.ResponseWriter) {
	fmt.Fprintln(w, "deleted - stubbed")
}

func createSyllabus(w http.ResponseWriter) {
	fmt.Fprintln(w, "create-stubbed")
}

func updateSyllabus(w http.ResponseWriter) {
	fmt.Fprintln(w, "update-stubbed")
}
