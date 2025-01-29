package files

import (
	"errors"
	"os"
)

type ReadFileOptions struct {
	create bool
}

func ReadFile(path string, opts ReadFileOptions) ([]byte, error) {
	file, err := os.ReadFile(path)

	if err != nil && opts.create {
		_, err := os.Create(path)

		if err != nil {
			return nil, err
		}

		return os.ReadFile(path)
	}

	return file, err
}

type CreateFileOptions struct{ truncate bool }

func CreateFile(path string, opts CreateFileOptions) (*os.File, error) {
	var file *os.File

	if _, err := os.Stat(path); os.IsNotExist(err) || opts.truncate {
		file, err = os.Create(path)

		return file, err
	}

	return nil, errors.New("cannot create file: " + path)
}
