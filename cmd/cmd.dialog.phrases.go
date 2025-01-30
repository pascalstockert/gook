package cmd

var CronAddPhrases = []Phrase{
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
