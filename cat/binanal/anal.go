package binanal

import (
	"debug/elf"
	"debug/macho"
	"debug/pe"
	"errors"
)

var NotValidPEFileError = errors.New("Not a valid PE file")
var NotValidELFFileError = errors.New("Not a valid ELF file")
var NotValidMACHOFileError = errors.New("Not a valid MACHO file")

type PESectionData struct {
	Name           string
	Size           int
	Offset         int
	End            int
	VirtualSize    int
	VirtualAddress int
}

type MACHOSectionData struct {
	Name   string
	Size   int
	Offset int
	End    int
}

type ELFSectionData struct {
	Name   string
	Size   int
	Offset int
	End    int
}

func PEAnal(input string) ([]PESectionData, []string, []string, error) {
	// An array of arrays for storing the section offsets
	var sectionData []PESectionData
	var symbolsArr []string

	// Check for executable type
	peFmt, err := pe.Open(input)
	if err != nil {
		return sectionData, []string{}, []string{}, NotValidPEFileError
	}
	defer peFmt.Close()

	// Extract sections
	sections := peFmt.Sections
	for k := range sections {
		sec := sections[k]
		secName := sec.Name
		secSize := sec.Size
		secOffset := sec.Offset + 1
		secEnd := secOffset + secSize - 1
		secVSize := sec.VirtualSize
		secVAddr := sec.VirtualAddress

		sd := PESectionData{
			Name:           secName,
			Size:           int(secSize),
			Offset:         int(secOffset),
			End:            int(secEnd),
			VirtualSize:    int(secVSize),
			VirtualAddress: int(secVAddr),
		}

		sectionData = append(sectionData, sd)
	}

	// Extract symbols
	numberOfSymbols := peFmt.NumberOfSymbols
	if numberOfSymbols > 0 {
		symbols := peFmt.Symbols

		for k := range symbols {
			sym := symbols[k]
			symName := sym.Name
			symbolsArr = append(symbolsArr, symName)
		}
	}

	// Get imported libraries
	libraries, err := peFmt.ImportedLibraries()
	if err != nil {
		return sectionData, []string{}, []string{}, err
	}

	return sectionData, libraries, symbolsArr, nil
}

func MACHOAnal(input string) ([]MACHOSectionData, []string, []string, error) {
	// An array of arrays for storing the section offsets
	var sectionData []MACHOSectionData

	// Check for executable type
	machoFmt, err := macho.Open(input)
	if err != nil {
		return sectionData, []string{}, []string{}, NotValidMACHOFileError
	}
	defer machoFmt.Close()

	// Extract sections
	sections := machoFmt.Sections
	for k := range sections {
		sec := sections[k]
		secName := sec.Name
		secSize := sec.Size
		secOffset := sec.Offset + 1
		secEnd := int(secOffset) + int(secSize) - 1

		sd := MACHOSectionData{
			Name:   secName,
			Size:   int(secSize),
			Offset: int(secOffset),
			End:    int(secEnd),
		}

		sectionData = append(sectionData, sd)
	}

	// Get imported symbols
	symbolsArr, err := machoFmt.ImportedLibraries()
	if err != nil {
		return sectionData, []string{}, []string{}, err
	}

	// Get imported libraries
	libraries, err := machoFmt.ImportedLibraries()
	if err != nil {
		return sectionData, []string{}, []string{}, err
	}

	return sectionData, libraries, symbolsArr, nil
}

func ELFAnal(input string) ([]ELFSectionData, []string, []string, error) {
	// An array of arrays for storing the section offsets
	var sectionData []ELFSectionData

	// Check for executable type
	elfFmt, err := elf.Open(input)
	if err != nil {
		return sectionData, []string{}, []string{}, NotValidELFFileError
	}
	defer elfFmt.Close()

	// Extract sections
	sections := elfFmt.Sections
	for k := range sections {
		sec := sections[k]
		secName := sec.Name
		secSize := sec.Size
		secOffset := sec.Offset + 1
		secEnd := int(secOffset) + int(secSize) - 1

		sd := ELFSectionData{
			Name:   secName,
			Size:   int(secSize),
			Offset: int(secOffset),
			End:    int(secEnd),
		}

		sectionData = append(sectionData, sd)
	}

	// Get imported symbols
	symbolsArr, err := elfFmt.ImportedLibraries()
	if err != nil {
		return sectionData, []string{}, []string{}, err
	}

	// Get imported libraries
	libraries, err := elfFmt.ImportedLibraries()
	if err != nil {
		return sectionData, []string{}, []string{}, err
	}

	return sectionData, libraries, symbolsArr, nil
}
