package pages

import (
	"SuBlog/models"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type PageData struct {
	Posts []models.Post `json:"Posts"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://127.0.0.1:8080/api/posts")
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

	var pageData PageData
	err = json.Unmarshal(body, &pageData)
	if err != nil {
		fmt.Println("Error marshall ", err)
		fmt.Println(resp.Body, string(body))
		return
	}

	tmpl, err := template.ParseFiles("\\routing\\routes\\pages\\index.html") // Replace with the path to your HTML template file
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}
