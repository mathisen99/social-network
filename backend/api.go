package backend

import (
	"database/sql"
	"fmt"
	"net/http"
)

func HandleSomeEndpoint(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Fprint(w, "Hello Gritlab!, this is the API endpoint!, you can use this endpoint to make requests to the backend.\n We can also use this endpoint to make requests to the database.")
}
