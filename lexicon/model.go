package lexicon

type lexicoModel struct {
	Id       int    `json:"lexicon_id"`
	Lexema   string `json:"lexema"`
	Canonica string `json:"canonica"`
	Class    string `json:"classe"`
	Forma    string `json:"forma"`
	Tempo    string `json:"tempo"`
	Tipo     string `json:"tipo"`
	Valor    string `json:"valor"`
	Flexao   string `json:"flexao"`
}

type flexaoModel struct {
	Lexico int    `json:"lexico_id"`
	Flexao string `json:"flexao"`
}
