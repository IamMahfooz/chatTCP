package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/user/register", app.registerUser)
	r.Post("/user/logout", app.userLogout)
	r.Post("/user/send", app.sendMessage)
	r.Post("/user/login", app.userLogin)
	r.Post("/user/onlineusers", app.onlineUsers)
	r.Post("/user/deleteuser", app.deleterUser)
}
