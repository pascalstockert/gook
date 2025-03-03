package cobra_cmd

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

type Phrase struct {
	Claim    string
	Options  []string
	Validate []func(string) error
}

func Dialog(reader bufio.Reader, phrases []Phrase) ([]string, error) {
	responses := make([]string, len(phrases))

	for index, phrase := range phrases {
		var response string
		var err error

		for {
			response, err = prompt(&reader, phrase)

			if err != nil {
				return nil, err
			}

			if phrase.Validate != nil {
				err = validatePhraseResponse(phrase, response)

				if err != nil {
					fmt.Println(err)

					continue
				}
			}

			responses[index] = response
			break
		}
	}

	return responses, nil
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

func validatePhraseResponse(phrase Phrase, response string) error {
	var errs []error

	for _, validateFn := range phrase.Validate {
		errs = append(errs, validateFn(response))
	}

	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}
