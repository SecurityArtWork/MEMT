package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	SERVER = "127.0.0.1:8888"
)

type ArtRes struct {
	Ssdeep    string   `json:"ssdeep"`
	Md5       string   `json:"md5"`
	Sha1      string   `json:"sha1"`
	Sha256    string   `json:"sha256"`
	Sha512    string   `json:"sha512"`
	Format    string   `json:"format"`
	Symbols   []string `json:"symbols"`
	Imports   []string `json:"imports"`
	Sections  []string `json:"sections"`
	Arch      string   `json:"arch"`
	Strain    string   `json:"strain"`
	Mutations []string `json:"mutations"`
	Siblings  []string `json:"siblings"`
}

type DataRes struct {
	Ecode int    `json:"ecode"`
	Msg   string `json:"msg"`
	Data  ArtRes `json:"data"`
}

type SearchRes struct {
	Ecode int    `json:"ecode"`
	Msg   string `json:"msg"`
	Goto  string `json:"goto"`
}

func main() {
	// Route definitions
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", pongerAPI).Methods("GET", "POST")
	router.HandleFunc("/api/v0/search/{hash}", searchAPI).Methods("GET")
	router.HandleFunc("/api/v0/malware/info/{hash}", infoAPI).Methods("GET")
	router.HandleFunc("/api/v0/malware/info/{hash}/{uuid}", infoUUIDAPI).Methods("GET")

	// Run server
	log.Printf("Running server on: http://%s", SERVER)
	log.Fatal(http.ListenAndServe(SERVER, router))
}

// =============
// = Endpoints =
// =============

// Ponger returns provides a basic echo server
// 		curl http://127.0.0.1:8888/
func pongerAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintln(w, "Pong")
		return
	}

	if r.Method == "POST" {
		type P struct {
			Msg string `json:"msg"`
		}

		res := P{"Pong"}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Println(err)
			sendJSONError(w, http.StatusInternalServerError, err)
		}
		return
	}
}

// Search endpoint api mocking
//		curl http://127.0.0.1:8888/api/v0/search/1234567890abcdef
//  	curl http://127.0.0.1:8888/api/v0/search/1234567890abcdef
func searchAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]

	if hash == "1234567890abcdef" {
		resURL := "http://" + SERVER + "/api/v0/malware/info/" + hash
		v := SearchRes{302, "Asset already analysed", resURL}
		json.NewEncoder(w).Encode(v)
		return
	} else {
		// resURL := "http://" + SERVER + "/api/v0/malware/info/" + hash + "/UUID-XXXX-YYYYYYYYY"
		// v := SearchRes{200, "Analysis has been launch in background", resURL}
		// json.NewEncoder(w).Encode(v)
		// return
		v := SearchRes{404, "This element does not exist", ""}
		json.NewEncoder(w).Encode(v)
		return
	}
}

// Info endpoint api mocking
//  	curl http://127.0.0.1:8888/api/v0/malware/info/1234567890abcdef
func infoAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]

	if hash == "1234567890abcdef" {
		a := ArtRes{
			Ssdeep:    "1234567890",
			Md5:       "1234567890",
			Sha1:      "1234567890",
			Sha256:    "1234567890",
			Sha512:    "1234567890",
			Format:    "pe",
			Symbols:   []string{"a", "b"},
			Imports:   []string{"a", "b"},
			Sections:  []string{"a", "b"},
			Arch:      "amd64",
			Strain:    "",
			Mutations: []string{"0987654321", "5647382910", "4536789013"},
			Siblings:  []string{""},
		}

		v := DataRes{
			302,
			"Asset already analysed",
			a,
		}

		json.NewEncoder(w).Encode(v)
		return
	} else {
		v := SearchRes{404, "This element does not exist", ""}
		json.NewEncoder(w).Encode(v)
		return
	}
}

// Info endpoint with uuid api mocking
func infoUUIDAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// hash := vars["hash"]
	uuid := vars["uuid"]

	if uuid == "UUID-XXXX-YYYYYYYYY" {
		a := ArtRes{
			Ssdeep:    "1234567890",
			Md5:       "1234567890",
			Sha1:      "1234567890",
			Sha256:    "1234567890",
			Sha512:    "1234567890",
			Format:    "pe",
			Symbols:   []string{"a", "b"},
			Imports:   []string{"a", "b"},
			Sections:  []string{"a", "b"},
			Arch:      "amd64",
			Strain:    "",
			Mutations: []string{"0987654321", "5647382910", "4536789013"},
			Siblings:  []string{""},
		}

		v := DataRes{
			302,
			"Asset already analysed",
			a,
		}

		json.NewEncoder(w).Encode(v)
		return
	} else {
		v := SearchRes{404, "This element does not exist", ""}
		json.NewEncoder(w).Encode(v)
		return
	}
}

// TODO: Upload endpoint
// resURL := "http://" + SERVER + "/api/v0/malware/info/" + hash + "/UUID-XXXX-YYYYYYYYY"
// v := SearchRes{200, "Analysis has been launch in background", resURL}
// json.NewEncoder(w).Encode(v)
// return

// ====================
// = Helper functions =
// ====================

func sendJSONError(rw http.ResponseWriter, code int, err error) {
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(code)
	log.Println(err)
	if err := json.NewEncoder(rw).Encode(err); err != nil {
		panic(err)
	}
}
