package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ApiData() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Asas")
		vars := mux.Vars(r)
		w.WriteHeader(http.StatusOK)
		fmt.Println("asasasassa")
		fmt.Fprintf(w, "Category: %v\n", vars["coin"])
	}
}
