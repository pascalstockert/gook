package main

import (
	"go-webhook/internal/cobra-cmd"
	"go-webhook/pkg/env"
)

func main() {
	env.Init()

	cobra_cmd.Execute()
}
