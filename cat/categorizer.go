package main

import (
	// "encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/dutchcoders/gossdeep"
	"github.com/hcninja/malpic/binanal"
	// "github.com/hcninja/malpic/image"
	// "github.com/securityartwork/hashing"
)

var thresholdFlag int
var dirFlag, imgoutFlag string
var verboseFlag bool

type Artifact struct {
	Ssdeep      string
	Md5         string
	Sha1        string
	Sha256      string
	Sha512      string
	Strain      string // if strain nil, else strain hash
	Format      string
	Symbols     []string
	Imports     []string
	Sections    []string
	Mutations   []string
	Photodir    string
	Artifactdir string
	// Arch        string
	// IP          string
}

func init() {
	flag.BoolVar(&verboseFlag, "verbose", false, "Goes verbose.")
	flag.StringVar(&dirFlag, "dir", "./", "Dir to scan.")
	flag.StringVar(&imgoutFlag, "imgout", "/tmp", "Output directory of generated pictures.")
	flag.IntVar(&thresholdFlag, "threshold", 1, "Sets threshold to compare")
}

func main() {
	flag.Parse()

	if thresholdFlag > 100 || thresholdFlag < 1 {
		fmt.Println("[!] Threshold can not be over 100% nor under 1%")
		os.Exit(1)
	}

	catalog()
}

// Runs the artifact cataloger
func catalog() {
	// Artifact array to store all the artifacts structs
	var artifactArray []Artifact
	var strains []string
	var mutations []string

	dir, err := ioutil.ReadDir(dirFlag)
	checkErr(err)

	dbgPrint("Calculating SSDEEP")
	// Calculate ssdeep for all the files and create Artifact
	for k := range dir {
		var element Artifact
		fileName := dir[k].Name()
		fileDir := path.Join(dirFlag, fileName)

		// Get DNA of binary
		hash, err := ssdeep.HashFilename(fileDir)
		checkErr(err)

		element.Ssdeep = hash
		element.Sha256 = fileName
		element.Artifactdir = dirFlag
		artifactArray = append(artifactArray, element)
	}

	// Generates the binary info
	dbgPrint("Artifact info extraction")
	for i := range artifactArray {
		fileDir := path.Join(artifactArray[i].Artifactdir, artifactArray[i].Sha256)

		// Get binary type
		// if err := binanal.IsELF(fileDir); err == nil {
		// 	dbgPrint("File is ELF")
		// 	artifactArray[i].Format = "elf"
		// } else if err := binanal.IsMACHO(fileDir); err == nil {
		// 	dbgPrint("File is MACH-O")
		// 	artifactArray[i].Format = "macho"
		// } else if err := binanal.IsPE(fileDir); err == nil {
		// 	dbgPrint("File is PE")
		// 	artifactArray[i].Format = "pe"
		// } else {
		// 	dbgPrint("File is not binary")
		// 	artifactArray[i].Format = "unknown"
		// }

		if err := binanal.IsPE(fileDir); err == nil {
			dbgPrint("File is PE")
			artifactArray[i].Format = "pe"
		}
	}

	// Genetic selector
	for i := range artifactArray {
		var mutsOfStrain []string
		atfNameA := artifactArray[i].Sha256
		isStrA := sliceContains(strains, atfNameA)
		isMutA := sliceContains(mutations, atfNameA)

		// Not a mutation nor strain, set as strain
		if !isStrA && !isMutA {
			strains = append(strains, atfNameA)
			artifactArray[i].Strain = ""
		}

		for j := range artifactArray {
			atfNameB := artifactArray[j].Sha256
			isSelf := atfNameA == atfNameB
			isStrB := sliceContains(strains, atfNameB)
			isMutB := sliceContains(mutations, atfNameB)

			// if A is B continue to next loop
			if isSelf {
				continue
			}

			perc, err := ssdeep.Compare(artifactArray[i].Ssdeep, artifactArray[j].Ssdeep)
			checkErr(err)

			if !isStrB && !isMutB && !isStrA && !isMutA {
				if perc >= thresholdFlag {
					mutsOfStrain = append(mutsOfStrain, atfNameB)
					mutations = append(mutations, atfNameB)
					artifactArray[j].Strain = atfNameA
				}
			}
		}

		// Append mutations of the strain to the strain
		artifactArray[i].Mutations = mutsOfStrain
	}

	// Vusual debug output
	if verboseFlag {
		for k := range artifactArray {
			fmt.Println("===========================")
			atf := artifactArray[k]
			fmt.Println("SSDEEP: " + atf.Ssdeep)
			fmt.Println("Format: " + atf.Format)
			fmt.Println("SHA: " + atf.Sha256)
			fmt.Println("Strain: " + atf.Strain)
			if len(atf.Mutations) > 0 {
				fmt.Println("Mutations: [")
				for _, mut := range atf.Mutations {
					fmt.Printf("\t%s,\n", mut)
				}
				fmt.Println("]")
			} else {
				fmt.Println("Mutations: []")
			}
		}
	}
}

// ===========
// = Helpers =
// ===========

// Checks if slice contains a given string
func sliceContains(list []string, str string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// Checks the error
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Print if debugFlag is set
func dbgPrint(msg string) {
	if verboseFlag {
		fmt.Println(fmt.Sprintf("[*] %s", msg))
	}
}

// Prints a nice msg
func msgPrint(msg ...string) {
	str := strings.Join(msg, " ")
	fmt.Println(str)
}
