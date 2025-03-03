package cobra_cmd

import (
	"bufio"
	"errors"
	"os"

	"github.com/robfig/cron"
	"go-webhook/pkg/helper"
)

func getCronAddDialogResponses() (
	name string,
	spec string,
	protocol string,
	location string,
) {
	phrases := []Phrase{
		{
			Claim: "Name the cronjob:",
		},
		{
			Claim:    "Specify the cron-spec:",
			Validate: []func(string) error{validateCronSpec},
		},
		{
			Claim: "Choose a protocol:",
			Options: []string{
				"http",
			},
		},
		{
			Claim: "Specify the resource location:",
		},
	}

	reader := getNewReader()
	responses, _ := Dialog(*reader, phrases)

	return helper.Destructure4(responses)
}

func validateCronSpec(response string) error {
	_, err := cron.ParseStandard(response)

	if err != nil {
		return errors.New("invalid cron-spec, please check your input")
	}

	return nil
}

func getNewReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}
