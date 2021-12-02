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

// access control and  CORS middleware
func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

var Lexicon *mux.Router //capital letter let this be exported
var Parser *mux.Router  //capital letter let this be exported

func main() {

	godotenv.Load(".env")
	lexicon.ConnectPsql()
	router := mux.NewRouter()
	router.Use(accessControlMiddleware)

	Lexicon = router.PathPrefix("/lexicon/").Subrouter()
	Parser = router.PathPrefix("/parser/").Subrouter()
	lexicon.CreateRoutes(Lexicon)
	parser.CreateRoutes(Parser)

	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8800",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal("Not connected: ", srv.ListenAndServe())

}

// "Access-Control-Allow-Origin", "*");
// resonse_object.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
