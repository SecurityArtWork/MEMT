package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/dutchcoders/gossdeep"
	"github.com/securityartwork/cat/binanal"
	"github.com/securityartwork/cat/image"
	// "github.com/securityartwork/cat/hashing"
)

var thresholdFlag int
var dirFlag, imgoutFlag string
var verboseFlag bool

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

	dbgPrint("Calculating SSDEEP.")
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
		element.ArtifactDir = dirFlag
		element.ImageDir = imgoutFlag
		artifactArray = append(artifactArray, element)
	}

	// Generates the binary info
	dbgPrint("Artifact info extraction.")
	for i := range artifactArray {
		fileDir := path.Join(artifactArray[i].ArtifactDir, artifactArray[i].Sha256)
		fullImageDir := path.Join(artifactArray[i].ImageDir, artifactArray[i].Sha256)

		// Reads the artifact into a binary array
		binaryArray, err := ioutil.ReadFile(fileDir)
		checkErr(err)

		if sectionData, libraries, symbols, err := binanal.PEAnal(fileDir); err == nil {
			// Check and extract data if PE
			err := setArtifactData(&artifactArray[i], "pe", symbols, libraries)
			checkErr(err)
			generateColorImage(fullImageDir, binaryArray, sectionData)
		} else if sectionData, libraries, symbols, err := binanal.ELFAnal(fileDir); err == nil {
			// Check and extract data if ELF
			err := setArtifactData(&artifactArray[i], "elf", symbols, libraries)
			checkErr(err)
			generateColorImage(fullImageDir, binaryArray, sectionData)
		} else if sectionData, libraries, symbols, err := binanal.MACHOAnal(fileDir); err == nil {
			// Check and extract data if Mach-O
			err := setArtifactData(&artifactArray[i], "macho", symbols, libraries)
			checkErr(err)
			generateColorImage(fullImageDir, binaryArray, sectionData)
		} else {
			// Not a PE, ELF nor MACH-O
			err := setArtifactData(&artifactArray[i], "unknown", symbols, libraries)
			checkErr(err)
			generateImage(fullImageDir, binaryArray)
		}
	}

	// Genetic selector
	dbgPrint("Genetic classification.")
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

	// Prints the formatted JSON
	fmt.Println("[")
	for k := range artifactArray {
		jsonBytes, _ := json.MarshalIndent(artifactArray[k], "", "\t")
		fmt.Println(string(jsonBytes) + ",")
	}
	fmt.Println("]")
}

// Sets the artifact fields
func setArtifactData(artifact *Artifact, format string, symbols, libraries []string) error {
	// Format fields
	artifact.Format = format
	artifact.Symbols = symbols
	artifact.Imports = libraries

	// Read file for hashing
	fName := path.Join(artifact.ArtifactDir, artifact.Sha256)
	file, err := ioutil.ReadFile(fName)
	if err != nil {
		return err
	}

	// Copies file to md5 hash array
	md5hasher := md5.New()
	md5hasher.Write(file)
	md5hash := md5hasher.Sum(nil)

	// Copies file to sha1 hash array
	sha1hasher := sha1.New()
	sha1hasher.Write(file)
	sha1hash := sha1hasher.Sum(nil)

	// Copies file to sha512 hash array
	sha512hasher := sha512.New()
	sha512hasher.Write(file)
	sha512hash := sha512hasher.Sum(nil)

	artifact.Md5 = fmt.Sprintf("%x", md5hash)
	artifact.Sha1 = fmt.Sprintf("%x", sha1hash)
	artifact.Sha512 = fmt.Sprintf("%x", sha512hash)

	return nil
}

// Encodes the binary in a colorful or B/W image
func generateColorImage(imgout string, binaryArray []byte, sectionData []binanal.SectionData) error {
	encoder, binImage := image.EncodeColor(binaryArray, sectionData)

	// Write image to file
	malPict, err := os.Create(imgout + ".png")
	if err != nil {
		return err
	}
	encoder.Encode(malPict, binImage)

	return nil
}

// Generates a B/W image file
func generateImage(imgout string, binaryArray []byte) error {
	// Encodes the binary in a colorful or B/W image
	encoder, binImage := image.EncodeBW(binaryArray)

	// Write image to file
	malPict, err := os.Create(imgout + ".png")
	if err != nil {
		return err
	}
	encoder.Encode(malPict, binImage)

	return nil
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
