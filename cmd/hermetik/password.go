package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"golang.org/x/term"
	"os"
)

func passwordInput(prompt string) string {
	fmt.Print(prompt)
	input, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read user input")
	}
	fmt.Println()
	return string(input)
}
