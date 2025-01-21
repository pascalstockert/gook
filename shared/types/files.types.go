package types

type FileParser struct {
	ParseEntries func(path string) []CronEntry
	AddEntry     func(entry CronEntry)
	RemoveEntry  func(id string)
	FileSuffix   string
}
