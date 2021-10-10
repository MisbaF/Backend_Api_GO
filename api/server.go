package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type User struct{
	Id string`json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	posts []Post
}
type Post struct{
	Id string `json:"id"`
	Caption string`json:"caption"`
	ImageURL string`json:"imageURL"`
	PostedTimestamp string `json:"postedTimestamp"`
}
type Server struct{
	users []User
	posts []Post
	Mux *http.ServeMux // based on what ur asking it routes

}
func NewServer() *Server{
	return &Server {
		users: []User{},
		Mux: http.NewServeMux(),

	}

}
func (s *Server) CreateUser() http.HandlerFunc {
	return func (w http.ResponseWriter,r *http.Request){
		if r.Method != "POST"{
			http.Error(w, "Method is not supported.", http.StatusBadRequest)
			return
		}
		var user User
		if err := json.NewDecoder( r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return

		}
		s.users = append(s.users, user)
		fmt.Println(user)
		//_, err := w.Write([]byte(rune(http.StatusOK)))
		//if err != nil {
		//	return
		//}

	}
}
func (s *Server) GetUserUsingId() http.HandlerFunc{
	var userIdExp = regexp.MustCompile(`/users/(?P<id>\d+)`)
	return func (w http.ResponseWriter, r *http.Request){
		if r.Method != "GET"{
			http.Error(w, "Method is not supported.", http.StatusBadRequest)
			return
		}
		var userId string

		match := userIdExp.FindStringSubmatch(r.URL.Path)
		if len(match) > 0 {
			result := make(map[string]string)
			for i, name := range userIdExp.SubexpNames() {
				if i != 0 && name != "" {
					result[name] = match[i]
				}
			}
			userId= result["id"]
		}

		for _, u := range s.users {
			if u.Id == userId {
				if err := json.NewEncoder( w).Encode(u); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return

				}
				return
			}
		}
		http.Error(w, "user not found", http.StatusInternalServerError)
	}
}

func (s *Server) CreatePost() http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		if r.Method != "POST"{
			http.Error(w, "Method is not supported.", http.StatusBadRequest)
			return
		}
		var post Post
		if err := json.NewDecoder( r.Body).Decode(&post); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return

		}
		s.posts = append(s.posts, post)
		fmt.Println(post)
	}
}
