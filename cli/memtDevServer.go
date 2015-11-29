package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type ResponseJSON struct {
	Ssdeep    string
	Md5       string
	Sha1      string
	Sha256    string
	Sha512    string
	Format    string
	Symbols   []string
	Imports   []string
	Sections  []string
	Arch      string
	Strain    string
	Mutations []string
	Siblings  []string
}

func main() {
	http.HandleFunc("/", testEndpoint)

	// Serve API
	log.Printf("Server running on: http://%s", "127.0.0.1:8888")
	err := http.ListenAndServe(socket, nil)
	if err != nil {
		return err
	}
}

// Analysis endpoint
func testEndpoint(w http.ResponseWriter, r *http.Request) {
	// var res ResponseJSON

	// Read full request body and limit it to 1MB
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024*1024))
	if err != nil {
		log.Fatal(err)
		return
	}

	// Close body on full read
	if err := r.Body.Close(); err != nil {
		log.Fatal(err)
		return
	}

	// Unmarshal json, and, if error return unprocessable entity
	if err := json.Unmarshal(body, &at); err != nil {
		sendJSONError(w, 422, err)
		return
	}
}
