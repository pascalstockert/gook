package files

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"go-webhook/pkg/types"
)

var suffix = ".json"

func GetJsonParser() *types.FileParser {
	var parser = types.FileParser{
		ParseEntries: parseEntries,
		WriteEntries: writeEntries,
		GetFilePath:  getFilePath,
		FileSuffix:   suffix,
	}

	return &parser
}

func parseEntries(path string) []types.CronEntry {
	data, err := ReadFile(path, ReadFileOptions{create: true})
	if err != nil {
		panic(err)
	}

	if len(data) == 0 {
		fmt.Println("No file or entries found in path: " + path)
		fmt.Println("Creating a new file")
		return []types.CronEntry{}
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

func getFilePath(fileName string) (string, error) {
	executableLocation, err := os.Executable()
	if err != nil {
		return "", err
	}

	pathArray := strings.Split(executableLocation, "/")
	workingDirectory := strings.Join(pathArray[:len(pathArray)-1], "/") + "/"

	fmt.Println(workingDirectory + fileName + suffix)

	return workingDirectory + fileName + suffix, nil
}
