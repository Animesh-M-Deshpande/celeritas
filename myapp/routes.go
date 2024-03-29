package main

import (
	"fmt"
	"myapp/data"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// middleware must come before any routes

	// add routes here
	a.get("/", a.Handlers.Home)
	a.App.Routes.Get("/go-page", a.Handlers.GoPage)
	a.App.Routes.Get("/jet-page", a.Handlers.JetPage)
	a.App.Routes.Get("/sessions", a.Handlers.SessionTest)

	a.App.Routes.Get("/users/login", a.Handlers.UserLogin)
	a.post("/users/login", a.Handlers.PostUserLogin)

	a.App.Routes.Get("/form", a.Handlers.Form)
	a.App.Routes.Post("/form", a.Handlers.PostForm)
	a.get("/json", a.Handlers.JSON)
	a.get("/xml", a.Handlers.XML)
	a.get("/download-file", a.Handlers.DownloadFile)

	a.get("/crypto", a.Handlers.TestCrypto)
	a.App.Routes.Get("/create-user", func(w http.ResponseWriter, r *http.Request) {
		u := data.User{
			FirstName: "Animesh",
			LastName:  "Deshpande",
			Email:     "animeshdeshpande@gmail.com",
			Active:    1,
			Password:  "testpassword",
		}

		id, err := a.Models.Users.Insert(u)

		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		fmt.Fprintf(w, "%d %s", id, u.FirstName)

	})

	a.App.Routes.Get("/get-all-users", func(w http.ResponseWriter, r *http.Request) {
		users, err := a.Models.Users.GetAll()

		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		for _, x := range users {
			fmt.Fprint(w, x.LastName)
		}
	})

	a.App.Routes.Get("/get-user/{id}", func(w http.ResponseWriter, r *http.Request) {

		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		u, _ := a.Models.Users.Get(id)

		fmt.Fprintf(w, "%s %s %s", u.FirstName, u.LastName, u.Email)
	})

	a.App.Routes.Get("/update-user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		u, err := a.Models.Users.Get(id)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		u.LastName = a.App.RandomString(10)

		validator := a.App.Validator(nil)
		u.LastName = ""

		u.Validate(validator)
		if !validator.Valid() {
			fmt.Fprint(w, "failed validatoin")
			return
		}
		err = u.Update(*u)

		if err != nil {
			a.App.ErrorLog.Println(err)
		}

		fmt.Fprintf(w, "updated last name to %s", u.LastName)

	})
	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
