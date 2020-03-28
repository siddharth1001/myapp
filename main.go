package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	// Declare the static file directory and point it to the
	// directory we just made
	staticFileDirectory := http.Dir("./assets/")

	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/birds/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/birds").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")

	r.HandleFunc("/", rootResponseHandler).Methods("GET")
	return r
}

func main() {
	fmt.Println("Starting server...")
	var connString = "root:@/bird_encyclopedia"
	if os.Getenv("CONNSTRING") != "" {
		connString = os.Getenv("CONNSTRING")
		fmt.Println("connString is ", connString)
		// "myuser:root123@tcp(mysql-db:3306)/bird_encyclopedia"
	} else {
		fmt.Println("NO connString. Connect to local")
	}

	db, err := sql.Open("mysql", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		// TODO: Getting connection refused 70% of times. Need to find the root cause!
		// Even root:root@tcp(mysql-db:3306)/bird_encyclopedia is not working
		fmt.Println("========= PING ERROR!")
		panic(err)
	}

	InitStore(&dbStore{db: db})

	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := newRouter()
	http.ListenAndServe(":8090", r)
}

func rootResponseHandler(w http.ResponseWriter, r *http.Request) {
	var responseString = `
				{
					"app name": "Bird Encyclopedia",
					"author": "Siddharth Rawat",
				}`
	_, _ = fmt.Fprintf(w, responseString)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there !!!!!")
}
