package parser

import (
	"encoding/json"
	"fmt"
)

type NODE struct {
	Class string  `json:"class"` //SN, SV, SP, SD...
	Value string  `json:"value"`
	Edges []*NODE `json:"children"` //{'DET','V'}
}

func createNode(class string, value string) *NODE {
	if value != "" {
		return &NODE{
			Class: class,
			Value: value,
		}
	}
	return &NODE{
		Class: class,
	}
}

func (n *NODE) appendNode(node *NODE) {
	n.Edges = append(n.Edges, node)
}

func (n *PARSER) printTree() []byte {
	r, err := json.MarshalIndent(n.tree, "", "	")
	if err != nil {
		fmt.Println("Err: ", err)
	}
	return r
}
