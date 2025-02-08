package cobra_cmd

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

type Phrase struct {
	Claim   string
	Options []string
}

func Dialog(reader bufio.Reader, phrases []Phrase) []string {
	responses := make([]string, len(phrases))

	for index, phrase := range phrases {
		response, e := prompt(&reader, phrase)

		if e != nil {
			panic(e)
		}

		responses[index] = response
	}

	return responses
}

func prompt(reader *bufio.Reader, phrase Phrase) (string, error) {
	if phrase.Options != nil {
		return promptSelect(phrase.Claim, phrase.Options)
	}

	return promptLine(reader, phrase.Claim)
}

func promptLine(reader *bufio.Reader, claim string) (string, error) {
	fmt.Println(claim)
	response, e := reader.ReadString('\n')

	if e != nil {
		return "", e
	}

	return strings.TrimSpace(response), nil
}

func promptSelect(claim string, options []string) (string, error) {
	response := ""
	prompt := &survey.Select{
		Message: claim,
		Options: options,
	}
	e := survey.AskOne(prompt, &response)

	if e != nil {
		return "", e
	}

	return response, nil
}
