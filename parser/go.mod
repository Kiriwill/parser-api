module github.com/kiriwill/parser-db-api/parser

replace github.com/kiriwill/parser-db-api/lexicon => ../lexicon

go 1.16

require (
	github.com/gorilla/mux v1.8.0
	github.com/kiriwill/parser-db-api/lexicon v0.0.0-00010101000000-000000000000
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac
)
