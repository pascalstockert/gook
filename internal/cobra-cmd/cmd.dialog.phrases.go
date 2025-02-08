package cobra_cmd

import (
	"bufio"
	"os"

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
			Claim: "Specify the cron-spec",
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
	responses := Dialog(*reader, phrases)

	return helper.Destructure4(responses)
}

func getNewReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}
