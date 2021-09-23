package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kiriwill/parser-db-api/lexicon"
	"github.com/kiriwill/parser-db-api/parser"
)

var Lexicon *mux.Router //capital letter let this be exported
var Parser *mux.Router  //capital letter let this be exported

func main() {
	godotenv.Load(".env")
	lexicon.ConnectPsql()
	router := mux.NewRouter()

	Lexicon = router.PathPrefix("/lexicon/").Subrouter()
	Parser = router.PathPrefix("/parser/").Subrouter()
	lexicon.CreateRoutes(Lexicon)
	parser.CreateRoutes(Parser)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8800",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
