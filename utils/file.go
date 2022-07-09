package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func FileExists(path string) bool {
	// check if file exist
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	return false
}

func SavePrivateKey(path string, privateKey *rsa.PrivateKey) {
	// check if file exist
	_, err := os.Stat(path)
	if err == nil {
		// file exist
		var overwriteInput string
		fmt.Printf("File (%s) exists! Overwrite? [y/n] ", path)
		fmt.Scanln(&overwriteInput)
		if strings.ToLower(overwriteInput) == "n" {
			log.Fatalf("File (%s) exists! Exiting..", path)
		} else if strings.ToLower(overwriteInput) == "y" {
			log.Printf("Overwriting file (%s)..", path)
		} else {
			log.Fatalln("[error] invalid input!")
		}
	}

	file, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	var pemPrivateBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	err = pem.Encode(file, pemPrivateBlock)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("private key (%s) created successfully!", path)
}

func SaveCSR(path string, csr []byte) {
	// check if file exist
	_, err := os.Stat(path)
	if err == nil {
		// file exist
		var overwriteInput string
		fmt.Printf("File (%s) exists! Overwrite? [y/n] ", path)
		fmt.Scanln(&overwriteInput)
		if strings.ToLower(overwriteInput) == "n" {
			log.Fatalf("File (%s) exists! Exiting..", path)
		} else if strings.ToLower(overwriteInput) == "y" {
			log.Printf("Overwriting file (%s)..", path)
		} else {
			log.Fatalln("[error] invalid input!")
		}
	}

	file, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	var pemPrivateBlock = &pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: csr,
	}
	err = pem.Encode(file, pemPrivateBlock)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("csr (%s) created successfully!", path)
}

func ReadCSRFile(csrFile string) string {
	if FileExists(csrFile) {
		content, err := ioutil.ReadFile(csrFile)
		if err != nil {
			log.Fatalln(err)
		}
		return string(content)
	} else {
		log.Fatalf("[error] csr file (%s) does not exist!", csrFile)
	}

	return ""
}
