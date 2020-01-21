package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird

func getBirdHandler(w http.ResponseWriter, r *http.Request) {

	birds, err := store.GetBirds()

	// Everything else is the same as before
	birdListBytes, err := json.Marshal(birds)

	if err != nil {
		fmt.Println(fmt.Errorf("Error in getBirdHandler: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(birdListBytes)
}

func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	bird := Bird{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error in createBirdHandler: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	fmt.Println("Creating bird :: ", bird)

	// The only change we made here is to use the `CreateBird` method instead of
	// appending to the `bird` variable like we did earlier
	err = store.CreateBird(&bird)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("Added a bird:", bird)
	http.Redirect(w, r, "/birds/", http.StatusFound)
}
