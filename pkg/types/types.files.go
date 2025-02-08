package types

type FileParser struct {
	ParseEntries func(path string) []CronEntry
	WriteEntries func(path string, entries []CronEntry)
	FileSuffix   string
}
