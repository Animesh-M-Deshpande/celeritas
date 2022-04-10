package celeritas

import "os"

func (c *Celeritas) CreateDirIfNotExists(path string) error {
	const mode = 0755

	if _, error := os.Stat(path); os.IsNotExist(error) {
		err := os.Mkdir(path, mode)

		if err != nil {
			return err
		}
	}

	return nil
}
