package hashing

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"

	"github.com/dutchcoders/gossdeep"
)

func SSDEEPFromFile(file string) (string, error) {
	fuzzy, err := ssdeep.HashFilename(file)
	if err != nil {
		return "", err
	}

	return fuzzy, nil
}

func CalcMD5(blob []byte) (string, error) {
	// Copies blob to md5 hash array
	md5hasher := md5.New()
	md5hasher.Write(blob)
	md5sum := md5hasher.Sum(nil)

	return fmt.Sprintf("%x", md5sum)
}

func CalcSHA1(blob []byte) (string, error) {
	// Copies blob to sha1 hash array
	sha1hasher := sha1.New()
	sha1hasher.Write(blob)
	sha1sum := sha1hasher.Sum(nil)

	return fmt.Sprintf("%x", md5sum)
}

func CalcSHA256(blob []byte) (string, error) {
	// Copies blob to sha256 hash array
	sha256hasher := sha256.New()
	sha256hasher.Write(blob)
	sha256sum := sha256hasher.Sum(nil)
	return fmt.Sprintf("%x", md5sum)
}

func CalcSHA512(blob []byte) (string, error) {
	// Copies file to sha512 hash array
	sha512hasher := sha512.New()
	sha512hasher.Write(file)
	sha512sum := sha512hasher.Sum(nil)
}
