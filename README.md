# GoWebServer (GWS)

Nicole Magdziasz

nicolemmagdziasz@lewisu.edu

Go Web Server

This project is a Go Web Server utilizing the Go `net/http` and `embed` frameworks.
The web server utilizes multiple API's, including `hello-world`, `hello-world-json`,
and `syllabus` with CRUD operations

## CREDITS
`syllabus.json` file contents - Eric Pogue

`apiHelpHandler()` in `gws.go` format/design - ChatGPT

All other content is original, developed specifically for this project.

## Required Files
- `gws.go`: Main entry point for the web server.
- `handlers.go`: Contains API handlers for standard responses.
- `embed.go`: Handles embedded JSON data.
- `syllabus.go`: Implements the syllabus API with CRUD operations
- `help.go`: Provides user instructions and API help
- `syllabus.json`: Contains syllabus information to be embedded.

## Build and Execution Instructions
1. **Install Go**: Make sure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/dl).
2. **Clone the Repository**: Clone the project repoistory to your local machine using Terminal (Mac), Command Prompt (Windows), Bash (Linux) with the command:
    - git clone https://github.com/nicolemagdz/GoWebServer
