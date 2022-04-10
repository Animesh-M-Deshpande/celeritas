package main

import (
	"log"
	"os"

	"github.com/Animesh-M-Deshpande/celeritas"
)

func initApplication() *application {

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//init celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)

	if err != nil {
		log.Fatal(err)
	}

	return nil //temp
}
