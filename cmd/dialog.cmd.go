package cmd

import (
	"bufio"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"strings"
)

type Phrase struct {
	Claim   string
	Options []string
}

func Dialog(reader bufio.Reader, phrases []Phrase) []string {
	responses := make([]string, len(phrases))

	for index, phrase := range phrases {
		responses[index] = prompt(&reader, phrase)
	}

	return responses
}

func prompt(reader *bufio.Reader, phrase Phrase) string {
	if phrase.Options != nil {
		return promptSelect(phrase.Claim, phrase.Options)
	}

	return promptLine(reader, phrase.Claim)
}

func promptLine(reader *bufio.Reader, claim string) string {
	fmt.Println(claim)
	response, _ := reader.ReadString('\n')

	return strings.TrimSpace(response)
}

func promptSelect(claim string, options []string) string {
	response := ""
	prompt := &survey.Select{
		Message: claim,
		Options: options,
	}
	_ = survey.AskOne(prompt, &response)

	return response
}
