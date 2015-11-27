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

type SectionData struct {
	Name   string
	Size   int
	Offset int
	End    int
}

func PEAnal(input string) ([]SectionData, []string, []string, error) {
	// An array of arrays for storing the section offsets
	var sectionData []SectionData

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

		sd := SectionData{
			Name:   secName,
			Size:   int(secSize),
			Offset: int(secOffset),
			End:    int(secEnd),
		}

		sectionData = append(sectionData, sd)
	}

	// Get imported symbols
	symbolsArr, err := peFmt.ImportedSymbols()
	if err != nil {
		return sectionData, []string{}, []string{}, err
	}

	// Get imported libraries
	libraries, err := peFmt.ImportedLibraries()
	if err != nil {
		return sectionData, []string{}, []string{}, err
	}

	return sectionData, libraries, symbolsArr, nil
}

func MACHOAnal(input string) ([]SectionData, []string, []string, error) {
	// An array of arrays for storing the section offsets
	var sectionData []SectionData

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

		sd := SectionData{
			Name:   secName,
			Size:   int(secSize),
			Offset: int(secOffset),
			End:    int(secEnd),
		}

		sectionData = append(sectionData, sd)
	}

	// Get imported symbols
	symbolsArr, err := machoFmt.ImportedSymbols()
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

func ELFAnal(input string) ([]SectionData, []string, []string, error) {
	// An array of arrays for storing the section offsets
	var sectionData []SectionData
	var symbolsArr []string

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
		secEnd := secOffset + secSize - 1

		sd := SectionData{
			Name:   secName,
			Size:   int(secSize),
			Offset: int(secOffset),
			End:    int(secEnd),
		}

		sectionData = append(sectionData, sd)
	}

	// Get imported symbols
	symbols, err := elfFmt.ImportedSymbols()
	if err != nil {
		return sectionData, []string{}, []string{}, err
	}

	if len(symbols) > 0 {
		for k := range symbols {
			symbolsArr = append(symbolsArr, symbols[k].Name)
		}
	}

	// Get imported libraries
	libraries, err := elfFmt.ImportedLibraries()
	if err != nil {
		return sectionData, []string{}, []string{}, err
	}

	return sectionData, libraries, symbolsArr, nil
}
