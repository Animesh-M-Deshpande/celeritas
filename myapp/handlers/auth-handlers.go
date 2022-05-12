package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handlers) UserLogin(w http.ResponseWriter, r *http.Request) {

	err := h.App.Render.Page(w, r, "login", nil, nil)

	if err != nil {
		h.App.ErrorLog.Println(err)
		return
	}

}

func (h *Handlers) PostUserLogin(w http.ResponseWriter, r *http.Request) {

	fmt.Println("inside post user login")
	err := r.ParseForm()

	if err != nil {
		fmt.Println("error1")
		w.Write([]byte(err.Error()))
		return
	}
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	fmt.Printf("forma data username %s, password: %s", email, password)
	user, err := h.Models.Users.GetByEmail(email)

	if err != nil {
		fmt.Println("error2")
		w.Write([]byte(err.Error()))
		return
	}

	matches, err := user.PasswordMatches(password)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte("error validating the password"))

		return
	}

	if !matches {
		w.Write([]byte("Invalid password"))
		return
	}

	h.App.Session.Put(r.Context(), "userID", user.ID)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
