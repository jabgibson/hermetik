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
		flagCipher  string
		flagEncrypt bool
		flagFile    string
		flagNew     bool
		flagSize    int
		flagVersion bool

		flagDecrypt bool
	)

	flag.BoolVar(&flagEncrypt, "enc", false, "encrypt file")
	flag.StringVar(&flagFile, "file", "", "file to encrypt")
	flag.StringVar(&flagCipher, "cipher", "cipher.hx", "cipher file")
	flag.BoolVar(&flagNew, "new", false, "create a new shift key")
	flag.IntVar(&flagSize, "size", 512, "size of a new shift key [defaults to random 512 - 1024")
	flag.BoolVar(&flagVersion, "v", false, "show version")

	flag.BoolVar(&flagDecrypt, "dec", false, "decrypt file")
	flag.Parse()

	if flagVersion {
		fmt.Printf("hermetik version %s\n", hermetik.Version)
		os.Exit(0)
	}

	if flagNew {
		// TODO make random size (right now defaulting to 512)
		secret := echoOffUserInput("input secret: ")
		log.Info().Msg("obtained user input secret")
		seed := hermetik.HashSeedFromString(secret)
		cipher := hermetik.BuildCipher(seed, flagSize)
		fmt.Print(string(cipher))
		//if err := os.WriteFile(flagCipher, cipher, 0600); err != nil {
		//	log.Fatal().Err(err).Msg("failed to write cipher")
		//}
		os.Exit(0)
	}

	if flagEncrypt {
		if flagCipher == "" {
			log.Fatal().Msg("must specify cipher")
		}
		if flagFile == "" {
			log.Fatal().Msg("must specify file to encrypt")
		}

		subjectBytes, err := os.ReadFile(flagFile)
		if err != nil {
			log.Fatal().Err(err).Str("file", flagFile).Msg("failed to read file")
		}
		cipherBytes, err := os.ReadFile(flagCipher)
		if err != nil {
			log.Fatal().Err(err).Str("file", flagCipher).Msg("failed to read cipher")
		}
		encryptedBytes, err := hermetik.Encrypt(cipherBytes, subjectBytes)
		if err != nil {
			log.Fatal().Err(err).Str("subject", flagFile).Str("cipher", flagCipher).Msg("failed to encrypt")
		}

		fmt.Print(string(encryptedBytes))
		os.Exit(0)
	}

	if flagDecrypt {
		if flagCipher == "" {
			log.Fatal().Msg("must specify cipher")
		}
		if flagFile == "" {
			log.Fatal().Msg("must specify file to encrypt")
		}
		subjectBytes, err := os.ReadFile(flagFile)
		if err != nil {
			log.Fatal().Err(err).Str("file", flagFile).Msg("failed to read file")
		}
		cipherBytes, err := os.ReadFile(flagCipher)
		if err != nil {
			log.Fatal().Err(err).Str("file", flagCipher).Msg("failed to read cipher")
		}
		decryptedBytes, err := hermetik.Decrypt(cipherBytes, subjectBytes)
		if err != nil {
			log.Fatal().Err(err).Str("subject", flagFile).Str("cipher", flagCipher).Msg("failed to decrypt")
		}

		fmt.Print(string(decryptedBytes))
		os.Exit(0)
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
	//	encrypted := svc.encrypt(fbytes)
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
	//	decrypted := svc.decrypt(fbytes)
	//	if flagOutFile == "" {
	//		flagOutFile = "decrypted-" + flagShiftKey
	//	}
	//	err = os.WriteFile(flagOutFile, decrypted, 0600)
	//	if err != nil {
	//		log.Fatal().Err(err).Msg("Failed to write file")
	//	}
	//}
}
