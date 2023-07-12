package routing

import (
	"SuBlog/routing/routes/pages"
	"SuBlog/routing/routes/posts"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func Init(database *mongo.Database) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/posts", func(w http.ResponseWriter, r *http.Request) {
		posts.All(w, r, database)
	}).Methods("GET")
	r.HandleFunc("/api/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		posts.Post(w, r, database)
	}).Methods("GET")
	r.HandleFunc("/api/posts/create", func(w http.ResponseWriter, r *http.Request) {
		posts.Create(w, r, database)
	}).Methods("POST")
	r.HandleFunc("/api/posts/delete", func(w http.ResponseWriter, r *http.Request) {
		posts.Delete(w, r, database)
	}).Methods("DELETE")

	r.HandleFunc("/post/{id}", pages.Post).Methods("GET")
	r.HandleFunc("/create", pages.Create).Methods("GET")
	r.HandleFunc("/", pages.Index).Methods("GET")

	return r
}
