package main

import (
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/securityartwork/cli/api"
)

var usage = `



	MEMT-CLI allows you to interact with the MEMT server API from the system
	console.

	WARNING: MEMT-CLI sends the artifact object to the server for an
	exhaustive analysis, if you don't like to do so, you can disable this
	behavior with the -nosend flag.



	MEMT-CLI has two main operating modes:

	artifact analysis mode: This mode is triggered with the -file flag, it will
	send the artifact to MEMT server and analyze the binary returning all the
	available information about this artifact.

	directory watch mode: This mode is triggered with the -watch flag and
	setting the watch mode with the -path flag. In this operation mode MEMT-CLI
	will watch all the new files dropped in the designed folder and upload them
	to MEMT server, returning detailed information about the artifact in case
	of being malware.



	Information returned by the server:

	Ssdeep
	Md5
	Sha1
	Sha256
	Sha512
	Arch
	Format
	Symbols
	Imports
	Sections
	Strain
	Mutations
	Siblings

`

const (
	VERSION = "0.0.1β"
)

var verboseFlag, watchFlag, infoFlag, noSendFlag bool
var fileFlag, pathFlag string

func init() {
	flag.StringVar(&fileFlag, "file", "", "File to scan.")
	flag.StringVar(&pathFlag, "path", "", "Path to watch.")
	flag.BoolVar(&watchFlag, "watch", false, "Watch a given path.")
	flag.BoolVar(&verboseFlag, "verbose", false, "Goes verbose.")
	flag.BoolVar(&infoFlag, "info", false, "Shows malhive info.")
	flag.BoolVar(&noSendFlag, "nosend", false, "Do not sent the artifact to memt server.")
}

func main() {
	flag.Parse()

	if infoFlag {
		fmt.Printf("MalHive CLI – v%s\n", VERSION)
		fmt.Println("\tMassive Early Malware Triage Command Line Interface.")
		fmt.Println("\tLicensed under GPLv2 – 2015")
		fmt.Println("\thttps://github.com/securityartwork/memt")
		fmt.Println("\t@bitsniper, @msanchez_87, @xumeiquer")
		fmt.Println(usage)
		os.Exit(0)
	}

	// Check for flag incompatibility
	if watchFlag && fileFlag != "" {
		alertPrint("-watch and -file flags are incompatible, use -path instead.")
		os.Exit(1)
	}

	if !watchFlag && fileFlag == "" {
		alertPrint("-file needs some file to work.")
		os.Exit(1)
	}

	// Send files SHA256
	if !watchFlag && fileFlag != "" {
		getInfoByFileHash(fileFlag)
		os.Exit(0)
	}

	// TODO: Send file to server

	// TODO: Send new files in path
	if watchFlag {
		alertPrint("Path monitoring NYI.")
	}
}

// ============
// = Handlers =
// ============

func getInfoByFileHash(fileDir string) {
	// Read file into byte array
	file, err := ioutil.ReadFile(fileDir)
	checkErr(err)

	// Hash file
	hashStr := hashFile(file)
	verbPrint("File hash: ", hashStr)
	verbPrint("Searching hash…")

	// Search if the hash is already analyzed
	searchRes, err := api.SearchHash(hashStr)
	// searchRes, err := api.SearchHash("1234567890abcdef")
	checkErr(err)

	res, err := json.MarshalIndent(searchRes, "", "\t")
	checkErr(err)
	verbPrint("Server response:\n", string(res))

	// check ecode
	switch searchRes.Ecode {
	case 302:
		verbPrint("Found.")
		msgPrint("Your analisys link is: ", searchRes.Goto)
		// Get analysis result
		getRes, err := api.GetInfo(searchRes.Goto)
		checkErr(err)

		msgPrint("Analysis result:")
		switch getRes.Ecode {
		case 302:
			// TODO: Refator to pretty printer function
			a := getRes.Data
			if a.Strain == "" {
				fmt.Printf("\tStrain: %s\n", "not a strain")
			} else {
				fmt.Printf("\tStrain: %s\n", a.Strain)
			}
			fmt.Printf("\tSSDEEP: %s\n", a.Ssdeep)
			fmt.Printf("\tMD5: %s\n", a.Md5)
			fmt.Printf("\tSHA1: %s\n", a.Sha1)
			fmt.Printf("\tSHA256: %s\n", a.Sha256)
			fmt.Printf("\tSHA512: %s\n", a.Sha512)
			fmt.Printf("\tFormat: %s\n", a.Format)
			fmt.Printf("\tArch: %s\n", a.Arch)
			fmt.Println("\tSymbols:")
			for k := range a.Symbols {
				fmt.Printf("\t\t %s\n", a.Symbols[k])
			}
			fmt.Println("\tImports:")
			for k := range a.Imports {
				fmt.Printf("\t\t %s\n", a.Imports[k])
			}
			fmt.Println("\tSections:")
			for k := range a.Sections {
				fmt.Printf("\t\t %s\n", a.Sections[k])
			}
			fmt.Println("\tMutations:")
			for k := range a.Mutations {
				fmt.Printf("\t\t %s\n", a.Mutations[k])
			}
			fmt.Println("\tSiblings:")
			for k := range a.Siblings {
				fmt.Printf("\t\t %s\n", a.Siblings[k])
			}

		case 404:
			alertPrint("Something weird happened, try it again later.")
		}

		return
	case 404:
		verbPrint("Not found.")
		if noSendFlag {
			msgPrint("Hash was not found, try to upload your sample without '-nosend' flag")
			return
		}

		verbPrint("Sending sample to MEMT")
		alertPrint("Not available yet…")
	}
}

// ===========
// = Helpers =
// ===========

// Calculates the hash of the given file
func hashFile(file []byte) string {
	// Hash file
	hasher := sha256.New()
	_, err := hasher.Write(file)
	checkErr(err)
	hash := hasher.Sum(nil)
	hashStr := fmt.Sprintf("%x", hash)

	return hashStr
}

func checkErr(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}

// Prints only debug messages
func verbPrint(msg ...string) {
	if verboseFlag {
		str := strings.Join(msg, " ")
		fmt.Println(fmt.Sprintf("[*] %s", str))
	}
}

// Prints alert messages with ! prefix
func alertPrint(msg ...string) {
	str := strings.Join(msg, " ")
	fmt.Println(fmt.Sprintf("[!] %s", str))
}

// Prints generic messages with + prefix
func msgPrint(msg ...string) {
	str := strings.Join(msg, " ")
	fmt.Println(fmt.Sprintf("[+] %s", str))
}
