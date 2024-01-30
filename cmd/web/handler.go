package main

import (
	"fmt"
	"net/http"
)

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user registration portal")
}
func (app *application) deleterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleting a user")
}
func(app *application) sendMessage(w http.ResponseWriter,r *http.Request){
	fmt.Println("used for sending message")
}
func(app *application)userLogin(w http.ResponseWriter,r *http.Request){
	fmt.Println("used when logging in")
}
func(app *application)onlineUsers(w http.ResponseWriter,r *http.Request){
	fmt.Println("lists all online users")
}
func(app *application)userLogout(w http.ResponseWriter,r *http.Request){
	fmt.Println("user gets log out")
}

