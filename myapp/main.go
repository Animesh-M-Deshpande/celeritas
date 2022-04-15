package main

import (
	"myapp/handlers"

	"github.com/Animesh-M-Deshpande/celeritas"
)

type application struct {
	App      *celeritas.Celeritas
	Handlers *handlers.Handlers
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
