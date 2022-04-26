package celeritas

import "net/http"

func (c *Celeritas) SessionLoad(next http.Handler) http.Handler {

	c.InfoLog.Println("session load called")
	return c.Session.LoadAndSave(next)
}
