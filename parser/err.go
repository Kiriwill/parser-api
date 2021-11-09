package parser

type DetailStr struct {
	Position    int     `json:"position"`
	Description string  `json:"description"`
	LastTree    []NODE  `json:"lasttree"`
	LastTokens  []TOKEN `json:"lasttokens"`
}

type ERR struct {
	Tpe    string    `json:"type"`
	Detail DetailStr `json:"detail"`
}
