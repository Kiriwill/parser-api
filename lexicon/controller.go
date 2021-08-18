package lexicon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getLexicon(res http.ResponseWriter, req *http.Request) {
	word := mux.Vars(req)
	var words []lexicoModel

	r := Db.queryWord(GetQuery, word["word"])
	for r.Next() {
		l := lexicoModel{}
		if err := r.Scan(&l.Id, &l.Lexema, &l.Canonica, &l.Class, &l.Forma, &l.Tempo, &l.Tipo, &l.Valor, &l.Flexao); err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		words = append(words, l)
	}
	if len(words) == 0 {
		http.Error(res, "Word not found.", http.StatusBadRequest)
		return
	}

	payload, err := json.MarshalIndent(words, "", "	")
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(payload)
}

func postLexicon(res http.ResponseWriter, req *http.Request) {
	var l lexicoModel
	l.Forma = "NULL"
	l.Tempo = "NULL"
	l.Tipo = "NULL"
	l.Valor = "NULL"

	err := json.NewDecoder(req.Body).Decode(&l)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	query := fmt.Sprintf(
		PostQuery, l.Lexema, l.Canonica, l.Canonica,
		l.Class, l.Tipo, l.Forma, l.Tempo, l.Valor)

	var id string
	err = Db.conn.QueryRow(query).Scan(&id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
}

func postFlexao(res http.ResponseWriter, req *http.Request) {
	var f flexaoModel

	err := json.NewDecoder(req.Body).Decode(&f)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if f.Flexao != "" && f.Lexico != 0 {
		query := fmt.Sprintf(FlexaoQuery, strconv.Itoa(f.Lexico), f.Flexao)

		_, err = Db.conn.Exec(query)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(res, "Bad Request", http.StatusBadRequest)
	}
	res.WriteHeader(http.StatusOK)

}

func deleteLexicon(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	query := fmt.Sprintf(
		DelLexiconQuery, vars["lexicon_id"])

	_, err := Db.conn.Exec(query)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusOK)
}
