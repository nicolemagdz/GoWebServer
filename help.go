package main

import "fmt"

func showHelp() {
	fmt.Println(`Go Web Server (gws) Application

Usage:
	go run *.go			-Start the Go Web Server
	go run *.go help		-Display help instructions
	
Available APIs:
	/hello-world			-Returns "Hello, World!" as plain text.
	/hello-world-json		-Returns a JSON object with a "Hello, World!" message.
	/hello-world-json-static	-Returns a static JSON message.
	/embedded-json			-Returns JSON data from an embedded file.
	
Ensure all required files are present and accessible.
	`)
}
