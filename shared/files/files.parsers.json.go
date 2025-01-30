package files

import (
	"encoding/json"

	"go-webhook/shared/types"
)

var suffix = ".json"

func GetJsonParser() *types.FileParser {
	var parser = types.FileParser{
		ParseEntries: parseEntries,
		WriteEntries: writeEntries,
		FileSuffix:   suffix,
	}

	return &parser
}

func parseEntries(path string) []types.CronEntry {
	data, err := ReadFile(path, ReadFileOptions{create: true})
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

func writeEntries(path string, entries []types.CronEntry) {
	data, err := json.Marshal(entries)
	if err != nil {
		panic(err)
	}

	file, err := CreateFile(path, struct{ truncate bool }{truncate: true})
	if err != nil {
		panic(err)
	}

	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}
}
