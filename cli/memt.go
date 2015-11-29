package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"os"
	"strings"

	// "github.com/securityartwork/Cli/api"
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

	// Send file SHA256
	if !watchFlag {
		// Open file
		file, err := os.Open(fileFlag)
		checkErr(err)
		defer file.Close()

		hash, err := sha256.(fileFlag)
		checkErr(err)
		verbPrint("{sha256:", fmt.Sprintf("%x}", hash))
		// TODO: Send to server
		os.Exit(0)
	}

	// TODO: Send new files in path
	if watchFlag {
		alertPrint("Path monitoring NYI.")
	}
}

// ===========
// = Helpers =
// ===========

func checkErr(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}

// Prints only debug messages
func verbPrint(msg ...string) {
	if verboseFlag {
		str := strings.Join(msg, " ")
		fmt.Println(fmt.Sprintf("[*] %s\n", str))
	}
}

// Prints alert messages with ! prefix
func alertPrint(msg ...string) {
	str := strings.Join(msg, " ")
	fmt.Println(fmt.Sprintf("[!] %s\n", str))
}

// Prints generic messages with + prefix
func msgPrint(msg ...string) {
	str := strings.Join(msg, " ")
	fmt.Println(fmt.Sprintf("[+] %s\n", str))
}
