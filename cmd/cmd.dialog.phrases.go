package cmd

import (
	"bufio"

	"go-webhook/shared/helper"
)

func getCronAddDialogResponses(reader bufio.Reader) (
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
	responses := Dialog(reader, phrases)

	return helper.Destructure4(responses)
}
