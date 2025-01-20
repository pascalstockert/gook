package env

import "os"

var environmentDefaults = []struct {
	key   string
	value string
}{
	{
		key:   "HTTP_SERVER_PID_FILE",
		value: "./http/http-server.pid",
	},
	{
		key:   "HTTP_SERVER_LOG_FILE",
		value: "./http/http-server.log",
	},
	{
		key:   "HTTP_SERVER_PORT",
		value: "4321",
	},
	{
		key:   "CRON_FILE_FORMAT",
		value: "json",
	},
}

func Init() {
	for _, env := range environmentDefaults {
		if envAlreadySet := os.Getenv(env.key); envAlreadySet == "" {
			_ = os.Setenv(env.key, env.value)
		}
	}
}
