package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-webhook/shared/env"
	files "go-webhook/shared/files/parsers"
)

type cobraCommandFunc func(cmd *cobra.Command, args []string)

var CronAdd cobraCommandFunc = func(cmd *cobra.Command, args []string) {
	var format, e = env.Get("CRON_FILE_FORMAT")
	if e != nil {
		fmt.Println(e)
		return
	}

	var parser, _ = files.GetParser(format)
	fmt.Println(parser.FileSuffix)
}
