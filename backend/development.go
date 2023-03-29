package backend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// ======================================================== ========================================================
//
// This file we only need during development, so we can run the Angular app and the Go backend at the same time
// When we are ready to deploy, we will build the Angular app and serve it from the Go backend (backend/server.go)
//
// ======================================================== ========================================================

// Global database variable that we can use in all the handlers
var db *sql.DB

func StartDevServer() {
	// Setting timeout, idle timeout, and read timeout we may need to change these values
	// these are used for the limiter
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  20 * time.Second,
	}

	// Multiplexer (What is a multiplexer? you ask...)
	// it makes sure to handle multiple requests at the same time. It works by receiving requests from the client,
	// distributing them to the appropriate services (in our case the handlers), and then sending the responses back to the client.
	// This allows the server to handle more requests with the same amount of resources, making it more efficient.
	mux := http.NewServeMux()

	// Handle the file servers for file upploads and the API endpoint (We may want diffrent directory for this)
	uploadFs := http.FileServer(http.Dir("./upload"))
	mux.Handle("/upload/", http.StripPrefix("/upload/", uploadFs))

	// Opening the databases for the backend (this is the only place we open the database)
	db = OpenDatabase()
	defer db.Close()

	// Create the proxy to the Angular app (We need this for development, so we dont have to build the Angular app every time we make a changes)
	angularProxy := createAngularProxy()
	mux.Handle("/", angularProxy)

	// Handle the API endpoints here...
	mux.HandleFunc("/api/", HandleMaker(HandleSomeEndpoint))

	// The Limit function is a rate limiter that limits the number of requests per second
	// We can use this to limit the number of requests per second to the API endpoints (Security reasons and good practice)
	limitedMux := Limit(mux)

	// Set up the server
	server.Addr = ":8080"
	server.Handler = limitedMux
	// Start the server
	fmt.Println("Server is listening on port 8080")
	log.Fatal(server.ListenAndServe())
}

// This creates the proxy to the Angular app (We need this for development, so we dont have to build the Angular app every time we make a changes)
func createAngularProxy() *httputil.ReverseProxy {
	target, _ := url.Parse("http://localhost:4200/")
	proxy := httputil.NewSingleHostReverseProxy(target)
	return proxy
}

// This helps us pass the database to the handlers so we don't have to open a new connection for each request
// Loopl at the example enpoint i put in we pass the db also not just the responsewriter and the request as we normally do
// In that example we pass the database to the handler so we can use it in the handler
func HandleMaker(fn func(http.ResponseWriter, *http.Request, *sql.DB)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, db)
	}
}
