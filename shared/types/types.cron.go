package types

type CronEntry struct {
	Id     string     `json:"id"`
	Name   string     `json:"name"`
	Spec   string     `json:"spec"`
	Action CronAction `json:"action"`
}

type CronAction struct {
	Type     CronActionType `json:"type"`
	Resource string         `json:"resource"`
}

type CronActionType string

const (
	CronActionTypeHttp CronActionType = "http"
)
