package files

import (
	"encoding/json"
	"go-webhook/shared/types"
)

var suffix = ".json"

func GetJsonParser() *types.FileParser {
	var parser = types.FileParser{
		ParseEntries: parseEntries,
		FileSuffix:   suffix,
	}

	return &parser
}

// TODO parse from file
func parseEntries(path string) []types.CronEntry {
	data, err := ReadFile(path, OpenOptions{create: true})
	if err != nil {
		panic(err)
	}

	var temp []types.CronEntry
	err = json.Unmarshal(data, &temp)
	if err != nil {
		panic(err)
	}

	return temp
}
