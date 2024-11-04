package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "help":
			showHelp()
			return
		default:
			fmt.Println("Unknown command. Use 'gws help' for instructions.")
			return
		}
	}

	http.HandleFunc("/hello-world", helloWorldHandler)
	http.HandleFunc("/hello-world-html", helloWorldHTMLHandler)
	http.HandleFunc("/hello-world-json", helloWorldJSONHandler)
	http.HandleFunc("/hello-world-json-static", helloWorldJSONStaticHandler)
	http.HandleFunc("/embedded-json", serveEmbeddedJSON)
	http.HandleFunc("/syllabus", syllabusHandler)
	http.HandleFunc("/help", apiHelpHandler)
	http.HandleFunc("/d6", d6Handler)

	fmt.Println("Starting server at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}

func apiHelpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `<h1>API Help</h1>
	<ul>
		<li><a href="/hello-world">/hello-world</a> - Returns "Hello World - GWS" as plain text.</li>
		<li><a href="/hello-world-html">/hello-world-html</a> - Returns "Hello World - GWS in HTML.</li>
		<li><a href="/hello-world-json">/hello-world-json</a> - Returns JSON {"message": "Hello World - GWS"}.</li>
		<li><a href="/hello-world-json-static">/hello-world-json-static</a> - Returns a static JSON message.</li>
		<li><a href="/embedded-json">/embedded-json</a> - Returns JSON data from an embedded file.</li>
		<li><a href="/syllabus">/syllabus</a> - Returns syllabus data.</li>
		<li><a href="/syllabus" onclick="alert('Use POST for Create')">/syllabus (POST)</a> - Create syllabus.</li>
		<li><a href="/syllabus" onclick="alert('Use PUT for Update')">/syllabus (PUT)</a> - Update syllabus.</li>
		<li><a href="/syllabus" onclick="alert('Use DELETE to delete syllabus')">/syllabus (DELETE)</a> - Delete syllabus.</li>
		<li><a href="/d6">/d6</a> - Returns a random integer 1-6 in JSON format.</li>
		<li><a href="/help">/help</a> - This help page.</li>
	</ul>`)
}

func d6Handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	roll := rand.Intn(6) + 1
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"roll": roll})
}