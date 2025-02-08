package files

import (
	"errors"

	"go-webhook/pkg/types"
)

func GetParser(format string) (*types.FileParser, error) {
	switch format {
	case "json":
		return GetJsonParser(), nil
	}

	return nil, errors.New("unknown parser format")
}
