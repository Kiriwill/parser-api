package parser

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func processSentence(s string) PARSER {
	text := s + " "
	sentence := strings.Split(strings.ToLower(text), " ")

	lexer := LEXER{input: sentence}
	lexer.lexemize()

	parser := PARSER{lexer: lexer}
	parser.nextToken()

	if !parser.IP() {
		parser.err = append(parser.err, ERR{
			Tpe: "parsing error",
			Detail: DetailStr{
				Position:    parser.lastPos - 1,
				Description: fmt.Sprintf("invalid sentence on position '%d'", parser.lastPos-1)},
		})
	}
	if len(lexer.err) != 0 || len(parser.err) != 0 {
		parser.err = append(parser.err, lexer.err...)
	}
	return parser
}

func postSentence(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	s, found := mux.Vars(req)["sentence"]
	if !found {
		http.Error(res, "Incorrect Sentence", http.StatusBadRequest)
		return
	}
	parser := processSentence(s)

	if len(parser.err) != 0 {
		payload, err := json.MarshalIndent(parser.err, "", "	")
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		res.WriteHeader(http.StatusBadRequest)
		res.Write(payload)
	} else if parser.lastPos == len(s) {
		payload, _ := json.MarshalIndent(ERR{
			Tpe: "string nao chegou até o final",
		}, "", "	")

		res.WriteHeader(http.StatusBadRequest)
		res.Write(payload)
	} else {
		payload, err := json.MarshalIndent(parser.tree, "", "	")
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		res.WriteHeader(http.StatusOK)
		res.Write(payload)
	}

}
