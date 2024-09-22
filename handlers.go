package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World - GWS!")
}

func helloWorldHTMLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<h1>Hello World &mdash; GWS</h1>")
}

func helloWorldJSONHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello World - GWS"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func helloWorldJSONStaticHandler(w http.ResponseWriter, r *http.Request) {
	staticJSON := `{"message": "This is a static JSON message"}`
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(staticJSON))
}
