package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"project/socialnetwork/post"
	"project/socialnetwork/server"
)

var srv *server.Server

func Routes(s *server.Server) {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/post", createPost)
	s.DB.AutoMigrate(&post.Post{})
	srv = s

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bonjour")
}

func createPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		p := post.Post{}
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		p.Create(srv)
	}

}
