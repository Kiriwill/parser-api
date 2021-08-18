package lexicon

import "github.com/gorilla/mux"

func CreateRoutes(r *mux.Router) {
	r.HandleFunc("/", postLexicon).Methods("POST")
	r.HandleFunc("/{lexicon_id}", deleteLexicon).Methods("DELETE")
	r.HandleFunc("/flexao", postFlexao).Methods("POST")
	r.HandleFunc("/{word}", getLexicon).Methods("GET")
}
