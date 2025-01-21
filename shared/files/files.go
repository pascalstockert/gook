package files

import "os"

type OpenOptions struct {
	create bool
}

func ReadFile(path string, opts struct{ create bool }) ([]byte, error) {
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
