package main

import (
	"InstagramBackendAPI/api"
	"net/http"
)

//Create a User
func main() {
	server := api.NewServer()
	mux := server.Mux
	mux.HandleFunc("/users",server.CreateUser())
	mux.HandleFunc("/posts",server.CreatePost())
	mux.HandleFunc("/users/",server.GetUserUsingId())
	http.ListenAndServe(":8080",mux)
}