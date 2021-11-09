package parser

type PARSER struct {
	lexer     LEXER
	tree      *NODE
	lastPos   int
	lastTrees []NODE
	err       ERR
}

type RESULT struct {
	Tree   *NODE   `json:"tree"`
	Tokens []TOKEN `json:"tokens"`
}

func (t *PARSER) nextNode(name string) *NODE {
	node := createNode(name, "") // crio atual
	t.tree.appendNode(node)      // incluo o atual
	t.tree = node                // aponta para o endereço do atual
	currentNode := t.tree        // guarda arvore atual - pq arvore pode mudar

	return currentNode
}

func (t *PARSER) backtrack(currentNode *NODE, current int, tokens []TOKEN) {
	// fmt.Println("backtrack de: ", *t.tree, " para: ", currentNode)
	//backtrack on Tree
	t.tree = currentNode

	t.tree.Edges = nil //reseta os filho s do nó atual para nil

	//backtrack on Token list
	t.lexer.currentPos = current
	t.lexer.tokens = tokens
	t.lexer.currentToken = t.lexer.tokens[t.lexer.currentPos]

}

func (t *PARSER) assertNotNullNode() {
	//backtrack on Tree
	for i, node := range t.tree.Edges {
		if len(node.Edges) == 0 {
			t.tree.Edges[i] = nil
		}
	}
	// t.tree.Edges
}

func (t *PARSER) nextToken() bool {
	t.lexer.currentPos += 1
	t.lexer.currentToken = t.lexer.tokens[t.lexer.currentPos]
	return true
}

func (t *PARSER) term(token string) bool {
	for _, v := range t.lexer.currentToken.Classe {
		if v == token {
			t.lexer.tokens[t.lexer.currentPos].Classe = []string{token}
			t.lexer.currentToken.Classe = []string{v}
			node := createNode(token, t.lexer.input[t.lexer.currentPos])
			t.tree.appendNode(node)

			t.nextToken()
			return true
		}
	}
	return false
}

func hasElementClass(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func hasElementClassStruct(s []NODE, str string) bool {
	for _, v := range s {
		if v.Class == str {
			return true
		}
	}
	return false
}

func (t *PARSER) keepTrack(lastNode *NODE) {
	sonHeigth := t.tree.Heigth + 1
	classes := []string{"AP", "AdvP", "DP", "NP", "NumP", "PossP", "PP", "QP", "VP"}
	if lastNode.Heigth < sonHeigth {
		lastNode.Heigth = sonHeigth
		if hasElementClass(classes, lastNode.Class) && !hasElementClassStruct(t.lastTrees, lastNode.Class) && lastNode.Edges != nil {
			t.lastTrees = append(t.lastTrees, *lastNode)
		}
	}

}
