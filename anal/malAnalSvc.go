package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	// "path"

	// "github.com/securityartwork/cat/binanal"
	// "github.com/securityartwork/cat/image"
	// "github.com/securityartwork/hashing"
)

const (
	VERSION = "0.0.1β"
)

var serviceFlag, infoFlag bool
var hostFlag, portFlag string

func init() {
	flag.BoolVar(&infoFlag, "-info", false, "Shows about info.")
	flag.BoolVar(&serviceFlag, "-daemon", false, "Analysis as μservice.")
	flag.StringVar(&hostFlag, "-host", "127.0.0.1", "Sets the μservice address.")
	flag.StringVar(&portFlag, "-port", "31337", "Sets the μservice port.")
}

// "ipMeta": [{"ip": "173.194.45.55",
//             "iso_code": "US",
//             "country": "United States",
//             "city": "Minneapolis",
//             "geo": [-93.2166, 44.9759],
//             "date": ""}],
//
// type IPMeta struct {
// 	IP      string    `json:"ip"`
// 	ISOCode string    `json:"iso_code"`
// 	Country string    `json:"country"`
// 	City    string    `json:"city"`
// 	Geo     []float32 `json:"geo"`
// 	Date    string    `json:"date"`
// }

type Artifact struct {
	Ssdeep      string   `json:"ssdeep"`
	Md5         string   `json:"md5"`
	Sha1        string   `json:"sha1"`
	Sha256      string   `json:"sha256"`
	Sha512      string   `json:"sha512"`
	Strain      string   `json:"strain"` // if strain nil, else strain hash
	Format      string   `json:"format"`
	Symbols     []string `json:"symbols"`
	Imports     []string `json:"imports"`
	Sections    []string `json:"sections"`
	Mutations   []string `json:"mutations"`
	ImageDir    string   `json:"imageDir"`
	ArtifactDir string   `json:"artifactDir"`
	Arch        string   `json:"arch"`
	Task        AnalTask `json:"taskData"`
}

type AnalTask struct {
	CeleryID string      `json:"celeryID"`
	MongoID  string      `json:"mongoID"`
	IPMeta   interface{} `json:"ipMeta"`
	Path     string      `json:"path"`
}

func main() {
	// Parse and check flags
	flag.Parse()

	// If service flag is set
	if serviceFlag {
		fmt.Println("[!] Service NYI")
		os.Exit(1)
	}

	// If info flag is set
	if infoFlag {
		fmt.Printf("MalHive analysis μservice v%s\n", VERSION)
		fmt.Println("\tLicensed under GPLv2 – 2015")
		fmt.Println("\thttps://github.com/securityartwork/memt")
		fmt.Println("\t@bitsniper, @msanchez_87, @xumeiquer")
	}

	if err := startServer(hostFlag, portFlag); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// RPC Server
func startServer(host string, port string) error {
	socket := host + ":" + port

	// Set / handler
	http.HandleFunc("/", analysisEndpoint)

	// Serve API
	log.Printf("Server running on: http://%s", socket)
	err := http.ListenAndServe(socket, nil)
	if err != nil {
		return err
	}

	return nil
}

// =======
// = API =
// =======

// Analysis endpoint
func analysisEndpoint(rw http.ResponseWriter, req *http.Request) {
	// New analysis task
	var at AnalTask
	var artifact Artifact

	// Read full request body and limit it to 1MB
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1024*1024))
	if err != nil {
		log.Fatal(err)
		return
	}

	// Close body on full read
	if err := req.Body.Close(); err != nil {
		log.Fatal(err)
		return
	}

	// Unmarshal json, and, if error return unprocessable entity
	if err := json.Unmarshal(body, &at); err != nil {
		sendJSONError(rw, 422, err)
		return
	}

	// Recover path of the artifact to analyze
	path := at.Path

	if err := generateHashes(&artifact, path); err != nil {
		sendJSONError(rw, http.StatusInternalServerError, err)
		return
	}

	if err := binaryData(&artifact, path); err != nil {
		sendJSONError(rw, http.StatusInternalServerError, err)
		return
	}

	artifact.Task = at

	// Send result of the analysis
	rw.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(artifact); err != nil {
		log.Println(err)
		sendJSONError(rw, http.StatusInternalServerError, err)
	}
}

// Generate the hashes of the binary file
func generateHashes(artifact *Artifact, path string) error {
	return nil
}

// Extracts the info from the binary
func binaryData(artifact *Artifact, path string) error {
	return nil
}

// Relocates the uploaded file
func relocate(artifact *Artifact, path string) error {
	return nil
}

// ===========
// = Helpers =
// ===========

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
