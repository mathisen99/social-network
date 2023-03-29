package backend

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// UNCOMMENT the DB and HandleMaker if you want to use a build version
// ALSO comment out the devlopment.go file they wont work together!!
// Mvh: Mathisen

/* // Global database variable that we can use in all the handlers
var db *sql.DB

// This helps us pass the database to the handlers so we don't have to open a new connection for each request
func HandleMaker(fn func(http.ResponseWriter, *http.Request, *sql.DB)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, db)
	}
} */

func StartServer() {
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

	// Serve the Angular app
	fs := http.FileServer(http.Dir("./build-version-0.1"))
	mux.Handle("/", fs)

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
