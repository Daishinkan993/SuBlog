package pages

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Create(w http.ResponseWriter, r *http.Request) {
	page, err := os.ReadFile("\\routing\\routes\\pages\\create.html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(page))
}
