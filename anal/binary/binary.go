package binary

import (
	// "debug/elf"
	// "debug/macho"
	"debug/pe"
	"errors"
)

var NotValidPEFileError = errors.New("Not a valid PE file")
var NotValidELFFileError = errors.New("Not a valid ELF file")
var NotValidMACHOFileError = errors.New("Not a valid MACHO file")

func PEAutopsy(fileDir string) error {
	// Open PE file
	peFile, err := pe.Open(name)
	if err != nil {
		return NotValidPEFileError
	}
	defer peFile.Close()

	// Extract all the valuable info
	architecture := peFile.Machine
	librariesArr, err := peFile.ImportedLibraries()
	if err != nil {
		return err
	}
	importedSymbols, err := peFile.ImportedSymbols()
	if err != nil {
		return err
	}
    peFile.
}
