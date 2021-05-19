package router

import (
	"me-english/packages/router/routes"
	"me-english/utils/resp"
	"net/http"

	"github.com/gorilla/mux"
)

func GoHomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp.Success(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "Welcome to MeEnglish GO API",
	})
}

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", GoHomePage)
	return routes.SetupRoutesWithMiddlewares(r)
}
