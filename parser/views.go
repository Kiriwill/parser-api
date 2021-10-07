package parser

import "github.com/gorilla/mux"

func CreateRoutes(r *mux.Router) {
	r.HandleFunc("/", postSentence).Methods("POST", "OPTIONS").Queries("sentence", "{sentence}").Headers()
}
