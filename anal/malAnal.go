package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/securityartwork/cat/binanal"
	"github.com/securityartwork/cat/image"
	"github.com/securityartwork/hashing"
)

const (
	VERSION = "0.0.1β"
)

var serviceFlag, infoFlag, toDBFlag bool
var hostFlag, portFlag, dbPortFlag, dbHostFlag, dirFlag, logFlag string

func init() {
	flag.BoolVar(&infoFlag, "-info", false, "Shows about info.")
	flag.BoolVar(&serviceFlag, "-daemon", false, "Analysis as μservice.")
	flag.StringVar(&hostFlag, "-host", "127.0.0.1", "Sets the μservice address.")
	flag.StringVar(&portFlag, "-port", "31337", "Sets the μservice port.")
}

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
}

func main() {
	// Parse and check flags
	flag.Parse()

	// If service flag is set
	if serviceFlag {
		fmt.Println("Service NYI")
		os.Exit(1)
	}

	// If info flag is set
	if infoFlag {
		fmt.Printf("MalHive analysis μservice v%s\n", VERSION)
		fmt.Println("\tLicensed under GPLv2 – 2015")
		fmt.Println("\thttps://github.com/securityartwork/memt")
		fmt.Println("\t@bitsniper, @msanchez_87, @xumeiquer")
	}

	pathInfo, err := ioutil.ReadDir(dirFlag)
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
