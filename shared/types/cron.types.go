package types

type CronEntry struct {
	Id      string
	Name    string
	Spec    string
	Closure func()
}
