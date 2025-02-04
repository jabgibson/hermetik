package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"golang.org/x/term"
	"os"
)

func echoOffUserInput(prompt string) string {
	_, err := fmt.Fprint(os.Stderr, prompt)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to write user input")
	}
	input, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read user input")
	}
	_, err = fmt.Fprintln(os.Stderr)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	return string(input)
}
