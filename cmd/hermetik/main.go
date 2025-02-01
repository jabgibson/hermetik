package main

import (
	"flag"
	"fmt"
	"github.com/jabgibson/hermetik"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	ll := 3
	envll := os.Getenv("H6k_LOG_LEVEL")
	if envll != "" {
		if l, err := strconv.Atoi(envll); err != nil {
			log.Fatal().Err(err).Msg("Invalid H6k_LOG_LEVEL environment variable")
		} else if l < -1 || l > 5 {
			log.Fatal().Err(err).Msg("Must -1 (trace) to 5 (panic)")
		} else {
			ll = l
		}
	}
	zerolog.SetGlobalLevel(zerolog.Level(ll))
}

func main() {
	var (
		flagPrompt  bool
		flagSeed    bool
		flagVersion bool

		flagNew      bool
		flagEncrypt  bool
		flagDecrypt  bool
		flagShiftKey string
		flagOutFile  string
	)

	flag.BoolVar(&flagPrompt, "prompt", false, "prompt for input")
	flag.BoolVar(&flagSeed, "seed", false, "generate a seed")
	flag.BoolVar(&flagVersion, "v", false, "show version")

	flag.BoolVar(&flagNew, "new", false, "create a new shift key")
	flag.BoolVar(&flagEncrypt, "enc", false, "encrypt file")
	flag.BoolVar(&flagDecrypt, "dec", false, "decrypt file")
	flag.StringVar(&flagShiftKey, "sk", "", "filename to encrypt/decrypt")
	flag.StringVar(&flagOutFile, "o", "", "filename to write encrypted/decrypted file")
	flag.Parse()

	if flagVersion {
		log.Printf("hermetik version %s\n", hermetik.Version)
		os.Exit(0)
	}

	if flagSeed {
		seed := echoOffUserInput("secret seed: ")
		log.Info().Msg("obtained user input seed")
		fmt.Println(seed)
		fmt.Println(hermetik.FNVSeedFromString(seed))
		fmt.Println(hermetik.HashSeedFromString(seed))
	}

	//if flagEncrypt && flagShiftKey != "" {
	//	password := echoOffUserInput("Enter password: ")
	//	fbytes, err := os.ReadFile(flagShiftKey)
	//	if err != nil {
	//		log.Fatal().Err(err).Msg("Failed to read file")
	//	}
	//	svc, err := hermetik.New(password, len(fbytes))
	//	if err != nil {
	//		log.Fatal().Err(err).Msg("Failed to create service")
	//	}
	//	encrypted := svc.Encrypt(fbytes)
	//	if flagOutFile == "" {
	//		flagOutFile = flagShiftKey + ".h6k"
	//	}
	//	err = os.WriteFile(flagOutFile, encrypted, 0600)
	//	if err != nil {
	//		log.Fatal().Err(err).Msg("Failed to write file")
	//	}
	//}
	//
	//if flagDecrypt && flagShiftKey != "" {
	//	password := echoOffUserInput("Enter password: ")
	//
	//	fbytes, err := os.ReadFile(flagShiftKey)
	//	if err != nil {
	//		log.Fatal().Err(err).Msg("Failed to read file")
	//	}
	//	svc, err := hermetik.New(password, len(fbytes))
	//	if err != nil {
	//		log.Fatal().Err(err).Msg("Failed to create service")
	//	}
	//	decrypted := svc.Decrypt(fbytes)
	//	if flagOutFile == "" {
	//		flagOutFile = "decrypted-" + flagShiftKey
	//	}
	//	err = os.WriteFile(flagOutFile, decrypted, 0600)
	//	if err != nil {
	//		log.Fatal().Err(err).Msg("Failed to write file")
	//	}
	//}
}
