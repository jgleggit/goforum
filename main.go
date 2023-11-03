package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"forum/dataaccess"
	"forum/router"
	"forum/handlers"
)

func main() {

	ConnectDatabase()
	StartRouter()
	StartWebServer()

}

func ConnectDatabase() {

	_, err := dataaccess.InitDatabase()
	if err != nil {
		log.Fatalf(`Failed to connect to "forum.db" database: %v\n`, err)
	}
	fmt.Printf(`Successfully connected to "forum.db" database`)
}

// var tmpl *template.Template
//
// var err error
// tmpl, err = template.ParseGlob("templates/*.html")
// if err != nil {
// 	log.Fatalf("Error: Failed to parse the template. %v", err) // log.Fatalf will log the error and call os.Exit(1)
// }

func StartRouter() {

	// Declare fileserver variable "fs" with
	// parameter "static" refering to the http directory
	fs := http.FileServer(http.Dir("static"))

	// Use http.NewServeMux() to create an empty multiplex route server.
	// The "mux" or router stores a mapping between urls paths
	// for the application and thier corresponding handlers.
	mux := http.NewServeMux()

	// Setup the routes multiplexer "mux" in the router package
	router.SetupRoutes(mux)

	// Strip the "/static/" prefix of any request to the static
	// files directory. Contains ["images", "styles", "js", "fonts"].
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	// Register the indexHandler function for the "/" endpoint.

	// Strips the "/uploads/" prefix from any request to the uploads
	// files directory. Contains images added to posts by user.
	/*
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
	*/
	mux.HandleFunc("/", handlers.IndexHandler)
}

func StartWebServer() {
	// Start a new web server listening on port 8000
	log.Print("Starting web server on IP and Port:  http://localhost:8000")
	webserver := http.ListenAndServe(":8000", mux)
	log.Print("Server started on port:8000")
	if webserver != nil {
		log.Fatal(webserver)
	}
}
