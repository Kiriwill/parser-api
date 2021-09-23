package parser

type DetailStr struct {
	Position    int    `json:"position"`
	Description string `json:"description"`
}

type ERR struct {
	Tpe    string    `json:"type"`
	Detail DetailStr `json:"detail"`
}
