package main

import(
	// Standard library packages
	"fmt"
	"net/http"
	
	// Third party packages
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()
	
	// Add a handler on /test
	r.GET("/", func(w http.ResponseWriter, http *http.Request, _ httprouter.Params) {
		// Simply write some test data for now
		fmt.Fprint(w, "Welcome!\n")
	})
	
	// Fire up the server
	http.ListenAndServe(":8080", r)
}