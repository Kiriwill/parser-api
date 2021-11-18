package parser

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
)

var Limiter = rate.NewLimiter(1, 3)

func assertNotNull(node *NODE) bool {
	var left bool
	if len(node.Edges) == 0 && node.Value == "" {
		return true
	}
	if len(node.Edges) == 1 {
		left = assertNotNull(node.Edges[0])
		if left {
			node.Edges = []*NODE{}
		}
	} else if len(node.Edges) == 2 {
		left = assertNotNull(node.Edges[0])
		right := assertNotNull(node.Edges[1])
		if left && right {
			node.Edges = []*NODE{}
		} else if left {
			node.Edges = []*NODE{node.Edges[1]}
		} else if right {
			node.Edges = []*NODE{node.Edges[0]}
		}
	}
	return false
}

func processSentence(s string) PARSER {
	text := s + " "
	sentence := strings.Split(strings.ToLower(text), " ")

	lexer := LEXER{input: sentence}
	lexer.lexemize()

	parser := PARSER{lexer: lexer}
	parser.nextToken()

	isSentenceValid := parser.root()
	gotTheEnd := (parser.lastPos != len(strings.Split(strings.ToLower(s), " "))+1)

	if !isSentenceValid || (isSentenceValid && !gotTheEnd && lexer.err.Tpe == "") {
		// se chegar no final mas não for valida mostra o que conseguiu
		parser.err = ERR{
			Tpe: "parsing",
			Detail: DetailStr{
				Description: strings.Join(lexer.input[parser.lastPos:], " "),
				LastTree:    parser.lastTrees,
				LastTokens:  parser.lexer.tokens,
			},
		}
	}
	if lexer.err.Tpe != "" {
		if isSentenceValid { //faltam palavras mas houve processamento de partes da sentença
			lexer.err.Detail.LastTree = append(parser.lastTrees, *parser.tree)
			lexer.err.Detail.LastTokens = parser.lexer.tokens
		} else { //não há o que mostrar porque nem processou
			lexer.err.Tpe = "critical"
		}
		parser.err = lexer.err
	}
	return parser
}

func postSentence(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	if !Limiter.Allow() {
		http.Error(res, http.StatusText(429), http.StatusTooManyRequests)
		return
	}

	s, found := mux.Vars(req)["sentence"]
	if !found {
		http.Error(res, "Incorrect Sentence", http.StatusBadRequest)
		return
	}
	parser := processSentence(s)

	if parser.err.Tpe != "" {
		payload, err := json.MarshalIndent(parser.err, "", "	")
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		res.WriteHeader(http.StatusBadRequest)
		res.Write(payload)
		// }
		// else if parser.lastPos != len(strings.Split(strings.ToLower(s), " "))+1 {
		// 	payload, _ := json.MarshalIndent(ERR{
		// 		Tpe: "complexity",
		// 		Detail: DetailStr{
		// 			Position:    parser.lastPos,
		// 			Description: "O sistema ainda não é capaz de processar sentenças complexas.",
		// 		}}, "", "	")

		// 	res.WriteHeader(http.StatusBadRequest)
		// res.Write(payload)
	} else {
		assertNotNull(parser.tree)
		result := RESULT{Tree: parser.tree, Tokens: parser.lexer.tokens}
		payload, err := json.MarshalIndent(result, "", "	")
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		res.WriteHeader(http.StatusOK)
		res.Write(payload)
	}

}
