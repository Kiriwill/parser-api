module github.com/kiriwill/parser-db-api

replace github.com/kiriwill/parser-db-api/lexicon => ./lexicon

replace github.com/kiriwill/parser-db-api/parser => ./parser

// +heroku goVersion go1.16
go 1.16

require (
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.3.0
	github.com/kiriwill/parser-db-api/lexicon v0.0.0-00010101000000-000000000000
	github.com/kiriwill/parser-db-api/parser v0.0.0-00010101000000-000000000000
)
