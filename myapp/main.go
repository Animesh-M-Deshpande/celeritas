package main

import "github.com/Animesh-M-Deshpande/celeritas"

type application struct {
	App *celeritas.Celeritas
}

func main() {

	c := initApplication()
	c.App.ListenAndServe()
}
