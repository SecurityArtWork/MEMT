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
	_, err := md5hasher.Write(blob)
	if err != nil {
		return "", err
	}
	md5sum := md5hasher.Sum(nil)

	return fmt.Sprintf("%x", md5sum), nil
}

func CalcSHA1(blob []byte) (string, error) {
	// Copies blob to sha1 hash array
	sha1hasher := sha1.New()
	_, err := sha1hasher.Write(blob)
	if err != nil {
		return "", err
	}
	sha1sum := sha1hasher.Sum(nil)

	return fmt.Sprintf("%x", sha1sum), nil
}

func CalcSHA256(blob []byte) (string, error) {
	// Copies blob to sha256 hash array
	sha256hasher := sha256.New()
	_, err := sha256hasher.Write(blob)
	if err != nil {
		return "", err
	}
	sha256sum := sha256hasher.Sum(nil)
	return fmt.Sprintf("%x", sha256sum), nil
}

func CalcSHA512(blob []byte) (string, error) {
	// Copies file to sha512 hash array
	sha512hasher := sha512.New()
	_, err := sha512hasher.Write(blob)
	if err != nil {
		return "", err
	}
	sha512sum := sha512hasher.Sum(nil)
	return fmt.Sprintf("%x", sha512sum), nil
}
