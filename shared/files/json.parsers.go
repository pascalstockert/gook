package files

import "go-webhook/shared/types"

func GetJsonParser() *types.FileParser {
	var parser = types.FileParser{
		ParseEntries: parseEntries,
		FileSuffix:   ".json",
	}

	return &parser
}

// TODO parse from file
func parseEntries(_ string) []types.CronEntry {
	return []types.CronEntry{
		{
			Id:   "1",
			Name: "pasu-home",
			Spec: "@every 5s",
			Action: types.CronAction{
				Type:     "http",
				Resource: "https://pasu.me",
			},
		},
	}
}
