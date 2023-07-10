package pages

import (
	"SuBlog/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://127.0.0.1:8080/api/posts/" + mux.Vars(r)["id"])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var post models.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	tmpl, err := template.ParseFiles("C:\\SuBlog\\routing\\routes\\pages\\post.html") // Replace with the path to your HTML template file
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = tmpl.Execute(w, post)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}
