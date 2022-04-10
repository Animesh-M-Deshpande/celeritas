package celeritas

import (
	"fmt"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Celeritas struct {
	AppName string
	Debug   bool
	Version string
}

func (c *Celeritas) New(rootPath string) error {

	pathConfig := initPaths{
		rootPath:   rootPath,
		folderName: []string{"handlers", "views", "migrations", "Temp", "Middlewares"},
	}

	err := c.Init(pathConfig)

	if err != nil {
		return err
	}

	err = c.checkDotEnv(rootPath)

	if err != nil {
		return err
	}

	//read  .env file
	err = godotenv.Load(rootPath + ".env")

	if err != nil {
		return err
	}

	return nil
}

func (c *Celeritas) Init(p initPaths) error {
	root := p.rootPath

	for _, path := range p.folderName {

		// create folder if that does not exists
		err := c.CreateDirIfNotExists(root + "/" + path)

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Celeritas) checkDotEnv(path string) error {
	err := c.CreateFileIfNotExist(fmt.Sprintf("%s/.env", path))

	if err != nil {
		return err
	}

	return nil
}
