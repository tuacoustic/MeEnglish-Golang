package routes

import (
	vocabulary "me-english/packages/vocabulary"
	"net/http"
)

var vocabularyRoutes = []Route{
	{
		Uri:     "/vocab/add",
		Method:  http.MethodPost,
		Handler: vocabulary.AddVocab,
	},
}
