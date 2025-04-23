package library

import (
	"log"
	"net/http"
)


type AppHelper func(w http.ResponseWriter, r *http.Request) (error) 



func CreateHandler(f AppHelper) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling request for:", r.URL.Path)
		if err := f(w, r); err != nil {
			NewTemplate().ExecuteTemplate(w, "error.html", map[string]any{
				"Title": "Error",
				"Error": "Something Went Wrong",
			})
		}
	}
}

