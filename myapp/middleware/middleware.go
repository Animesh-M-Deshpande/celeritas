package middleware

import (
	"myapp/data"

	"github.com/Animesh-M-Deshpande/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models *data.Models
}
