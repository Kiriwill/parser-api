package parser

import (
	"database/sql"
	"log"

	"github.com/kiriwill/parser-db-api/lexicon"
)

type TOKEN struct {
	valor  string
	classe []string
}

type LEXER struct {
	input        []string //sentenca de entrada
	currentPos   int
	currentToken TOKEN
	conn         *sql.DB
	tokens       []TOKEN
	err          []ERR
}

func (t *LEXER) createToken(classes []string) {
	elem := TOKEN{
		valor:  t.input[t.currentPos],
		classe: classes,
	}
	t.tokens = append(t.tokens, elem)
}

func (t *LEXER) getClasses(r *sql.Rows) []string {
	var classes []string
	var class string
	var kind string
	for r.Next() {
		if err := r.Scan(&class, &kind); err != nil {
			log.Fatal(err)
			t.err = append(t.err, ERR{
				Tpe:    "lexical error",
				Detail: DetailStr{Description: err.Error()},
			})
		}
		class = t.sintagmaToClass(class, kind)
		classes = append(classes, class)
	}
	if classes == nil && t.input[t.currentPos] != "" {
		t.err = append(t.err, ERR{
			Tpe:    "lexical error",
			Detail: DetailStr{Description: "word '%s' not founded."},
		})
	}
	return classes
}

func (t *LEXER) explodeContraction() {
	var words []string
	var word string
	r := lexicon.Db.QueryWord(lexicon.ConComQuery, t.input[t.currentPos])
	for r.Next() {
		if err := r.Scan(&word); err != nil {
			log.Fatal(err)
		}
		words = append(words, word)
	}
	t.insertWordOnPosAndOveride(t.currentPos, words[0], words[1])
}

func (t *LEXER) insertToken() bool {
	t.currentPos += 1
	rows := lexicon.Db.QueryWord(lexicon.BaseQuery, t.input[t.currentPos])
	classes := t.getClasses(rows)

	if classes != nil && classes[0] == "CON" {
		t.explodeContraction()
		return false
	}

	t.createToken(classes)
	return true
}

func (l *LEXER) initReinit() { l.currentPos = -1 }

func (l *LEXER) lexemize() {
	l.initReinit()
	for l.currentPos < len(l.input)-1 {
		if !l.insertToken() {
			l.currentPos -= 1
		}
	}
	l.initReinit()
}

func (l *LEXER) insertWordOnPosAndOveride(pos int, word1 string, word2 string) {
	//Insert token on array and overide at current position
	l.input = append(l.input, "")
	copy(l.input[pos+1:], l.input[pos:])
	l.input[pos] = word1
	l.input[pos+1] = word2
}

func (t *LEXER) sintagmaToClass(class string, kind string) string {
	if class == "PRO" {
		switch kind {
		case "Dem":
			return "D"
		case "Pes":
			return "D"
		case "Pos":
			return "POSS"
		case "Ind":
			return "Q"
		}
	}
	if class == "ART" {
		return "D"
	}

	return class
}

//fazer a query do CON para as duas palavras (pref/suf) e inserir nos tokens
