package main

import (
	"flag"
	"fmt"
	"github.com/jabgibson/hermetik"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	var (
		flagEncrypt  bool
		flagDecrypt  bool
		flagFilename string
		flagOutFile  string
		flagVersion  bool
		flagShiftKey string
	)
	flag.BoolVar(&flagEncrypt, "e", false, "encrypt file")
	flag.BoolVar(&flagDecrypt, "d", false, "decrypt file")
	flag.StringVar(&flagFilename, "f", "", "filename to encrypt/decrypt")
	flag.StringVar(&flagOutFile, "o", "", "filename to write encrypted/decrypted file")
	flag.BoolVar(&flagVersion, "v", false, "show version")
	flag.StringVar(&flagShiftKey, "sk", "", "shift key out")
	flag.Parse()

	if flagVersion {
		fmt.Printf("hermetik %s\n", hermetik.Version)
		os.Exit(0)
	}

	if flagEncrypt && flagFilename != "" {
		password := passwordInput("Enter password: ")
		fbytes, err := os.ReadFile(flagFilename)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to read file")
		}
		svc, err := hermetik.New(password, len(fbytes))
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
		password := passwordInput("Enter password: ")
		
		fbytes, err := os.ReadFile(flagFilename)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to read file")
		}
		svc, err := hermetik.New(password, len(fbytes))
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
