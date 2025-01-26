package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"go-webhook/shared/env"
	"go-webhook/shared/files"
	"os"
)

type cobraCommandFunc func(cmd *cobra.Command, args []string)

var CronAdd cobraCommandFunc = func(cmd *cobra.Command, args []string) {
	var format, e = env.Get("CRON_FILE_FORMAT")
	if e != nil {
		panic(e)
	}

	var parser, _ = files.GetParser(format)
	entries := parser.ParseEntries("./cron-entries" + parser.FileSuffix)
	fmt.Println(entries)

	reader := bufio.NewReader(os.Stdin)

	responses := Dialog(*reader, []Phrase{
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
	})

	fmt.Println(responses[0])
	fmt.Println(responses[1])
	fmt.Println(responses[2])
	fmt.Println(responses[3])
}
