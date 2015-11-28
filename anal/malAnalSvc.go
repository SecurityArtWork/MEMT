/*

// Test set

{
	"path": "~/Desktop/memtEnvEmu/uploads/wxp1.exe",
	"ipMeta": {
		"ip": "173.194.45.55",
	    "iso_code": "US",
	    "country": "United States",
	    "city": "Minneapolis",
	    "geo": [-93.2166, 44.9759],
	    "date": "1448670771"}
}


curl -s http://localhost:31337/ -d'{"path": "/Users/kajuryto/Desktop/memtEnvEmu/uploads/wxp1.exe", "ipMeta": {"ip": "173.194.45.55", "iso_code": "US", "country": "United States", "city": "Minneapolis", "geo": [-93.2166, 44.9759], "date":"1448670771"}}' |python -m json.tool

*/

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
	"path"

	"github.com/securityartwork/anal/hashing"
	"github.com/securityartwork/cat/binanal"
	"github.com/securityartwork/cat/image"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

const (
	VERSION = "0.0.1β"
)

var serviceFlag, infoFlag, toDBFlag bool
var hostFlag, portFlag, binDstFlag, imgDstFlag, dbHostFlag, dbPortFlag string

func init() {
	// Operation mode flags
	flag.BoolVar(&infoFlag, "info", false, "Shows about info.")
	flag.BoolVar(&toDBFlag, "todb", true, "Stores the info into db.")

	// API flags
	flag.StringVar(&dbHostFlag, "dbhost", "127.0.0.1", "Sets the mongodb address.")
	flag.StringVar(&dbPortFlag, "dbport", "27017", "Sets the mongodb port.")

	// API flags
	flag.StringVar(&hostFlag, "host", "127.0.0.1", "Sets the μservice address.")
	flag.StringVar(&portFlag, "port", "31337", "Sets the μservice port.")

	// Analysis flags
	flag.StringVar(&binDstFlag, "bindst", "/tmp/bin", "Binary final folder.")
	flag.StringVar(&imgDstFlag, "imgdst", "/tmp/img", "Images final folder.")
}

type Artifact struct {
	Ssdeep      string      `json:"ssdeep"`
	Md5         string      `json:"md5"`
	Sha1        string      `json:"sha1"`
	Sha256      string      `json:"sha256"`
	Sha512      string      `json:"sha512"`
	Strain      string      `json:"strain"` // if strain nil, else strain hash
	Format      string      `json:"format"`
	Symbols     []string    `json:"symbols"`
	Imports     []string    `json:"imports"`
	Sections    []string    `json:"sections"`
	Mutations   []string    `json:"mutations"`
	ImageDir    string      `json:"imageDir"`
	ArtifactDir string      `json:"artifactDir"`
	Arch        string      `json:"arch"`
	IPMeta      interface{} `json:"ipMeta"`
}

type AnalTask struct {
	IPMeta interface{} `json:"ipMeta"`
	Path   string      `json:"path"`
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
	filePath := at.Path

	// Generate hashes
	if err := generateHashes(&artifact, filePath); err != nil {
		sendJSONError(rw, http.StatusInternalServerError, err)
		return
	}

	// Set file pointers
	artifact.ArtifactDir = path.Join(binDstFlag, artifact.Sha256)
	artifact.ImageDir = path.Join(imgDstFlag, artifact.Sha256+".png")

	// Extract binary info
	if err := binaryData(&artifact, filePath); err != nil {
		sendJSONError(rw, http.StatusInternalServerError, err)
		return
	}

	// Move file to binary repository
	if err := relocate(&artifact, filePath); err != nil {
		sendJSONError(rw, http.StatusInternalServerError, err)
		return
	}

	// Sets request IP meta into artifact struct
	artifact.IPMeta = at.IPMeta

	// TODO: Catalog new sample

	// TODO: Insert artifact into DB
	if toDBFlag {
		fmt.Println("NYI")
	} else {
		// Debug send result of the analysis back
		rw.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(rw).Encode(artifact); err != nil {
			log.Println(err)
			sendJSONError(rw, http.StatusInternalServerError, err)
		}
	}
}

// ==============
// = binary ops =
// ==============
// Generate the hashes of the binary file
func generateHashes(artifact *Artifact, filePath string) error {
	// Read file to byte array
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	ssdeep, err := hashing.SSDEEPFromFile(filePath)
	if err != nil {
		return err
	}

	md5, err := hashing.CalcMD5(file)
	if err != nil {
		return err
	}

	sha1, err := hashing.CalcSHA1(file)
	if err != nil {
		return err
	}

	sha256, err := hashing.CalcSHA256(file)
	if err != nil {
		return err
	}

	sha512, err := hashing.CalcSHA512(file)
	if err != nil {
		return err
	}

	artifact.Ssdeep = ssdeep
	artifact.Md5 = md5
	artifact.Sha1 = sha1
	artifact.Sha256 = sha256
	artifact.Sha512 = sha512

	return nil
}

// Extracts the info from the binary
func binaryData(artifact *Artifact, filePath string) error {
	if sectionData, libraries, symbols, err := binanal.PEAnal(filePath); err == nil {
		artifact.Format = "pe"
		artifact.Imports = libraries
		artifact.Symbols = symbols
		artifact.Sections = extractSectionNames(sectionData)
		if err := generateColorImage(artifact.ImageDir, filePath, sectionData); err != nil {
			return err
		}
	} else if sectionData, libraries, symbols, err := binanal.ELFAnal(filePath); err == nil {
		artifact.Format = "elf"
		artifact.Imports = libraries
		artifact.Symbols = symbols
		artifact.Sections = extractSectionNames(sectionData)
		if err := generateColorImage(artifact.ImageDir, filePath, sectionData); err != nil {
			return err
		}
	} else if sectionData, libraries, symbols, err := binanal.MACHOAnal(filePath); err == nil {
		artifact.Format = "macho"
		artifact.Imports = libraries
		artifact.Symbols = symbols
		artifact.Sections = extractSectionNames(sectionData)
		if err := generateColorImage(artifact.ImageDir, filePath, sectionData); err != nil {
			return err
		}
	} else {
		artifact.Format = "unknown"
		artifact.Imports = libraries
		artifact.Symbols = symbols
		if err := generateImage(artifact.ImageDir, filePath); err != nil {
			return err
		}
	}

	return nil
}

// Relocates the uploaded file
func relocate(artifact *Artifact, filePath string) error {
	if err := os.Rename(filePath, artifact.ArtifactDir); err != nil {
		return err
	}

	return nil
}

// Encodes the binary in a colorful or B/W image
func generateColorImage(imgout, filePath string, sectionData []binanal.SectionData) error {
	// Read file to byte array
	binaryArray, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	encoder, binImage := image.EncodeColor(binaryArray, sectionData)

	// Write image to file
	malPict, err := os.Create(imgout)
	if err != nil {
		return err
	}
	encoder.Encode(malPict, binImage)

	return nil
}

// Generates a B/W image file
func generateImage(imgout, filePath string) error {
	// Read file to byte array
	binaryArray, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Encodes the binary in a colorful or B/W image
	encoder, binImage := image.EncodeBW(binaryArray)

	// Write image to file
	malPict, err := os.Create(imgout)
	if err != nil {
		return err
	}
	encoder.Encode(malPict, binImage)

	return nil
}

// ================
// = database ops =
// ================

// ===========
// = Helpers =
// ===========
func extractSectionNames(sectionData []binanal.SectionData) []string {
	var sectionNames []string
	for k := range sectionData {
		sectionNames = append(sectionNames, sectionData[k].Name)
	}

	return sectionNames
}

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
