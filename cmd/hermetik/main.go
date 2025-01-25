package main

import (
	"flag"
	"fmt"
	"github.com/jabgibson/hermetik/shift"
	"github.com/rs/zerolog/log"
	"golang.org/x/term"
	"os"
)

func main() {
	var (
		flagEncrypt  bool
		flagDecrypt  bool
		flagFilename string
		flagOutFile  string
	)
	flag.BoolVar(&flagEncrypt, "e", false, "encrypt file")
	flag.BoolVar(&flagDecrypt, "d", false, "decrypt file")
	flag.StringVar(&flagFilename, "f", "", "filename to encrypt/decrypt")
	flag.StringVar(&flagOutFile, "o", "", "filename to write encrypted/decrypted file")
	flag.Parse()

	fmt.Println("Enter Secret: ")
	input, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read password")
	}
	password := string(input)

	if flagEncrypt && flagFilename != "" {
		fbytes, err := os.ReadFile(flagFilename)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to read file")
		}
		svc, err := shift.New(password, len(fbytes))
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create service")
		}
		encrypted := svc.Encrypt(fbytes)
		if flagOutFile == "" {
			flagOutFile = flagFilename + ".h6k"
		}
		err = os.WriteFile(flagOutFile, encrypted, 0600)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to write file")
		}
	}

	if flagDecrypt && flagFilename != "" {
		fbytes, err := os.ReadFile(flagFilename)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to read file")
		}
		svc, err := shift.New(password, len(fbytes))
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create service")
		}
		decrypted := svc.Decrypt(fbytes)
		if flagOutFile == "" {
			flagOutFile = "decrypted-" + flagFilename
		}
		err = os.WriteFile(flagOutFile, decrypted, 0600)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to write file")
		}
	}
}
