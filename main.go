package main

import (
	"go-webhook/cmd"
	"go-webhook/shared/env"
)

func main() {
	env.Init()

	cmd.Execute()
}
