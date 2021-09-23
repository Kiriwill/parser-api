package parser

type PARSER struct {
	lexer   LEXER
	tree    *NODE
	lastPos int
	err     []ERR
}

func (t *PARSER) nextNode(name string) *NODE {
	node := createNode(name, "") // crio atual
	t.tree.appendNode(node)      // incluo o atual
	t.tree = node                // aponta para o endereço do atual
	currentNode := t.tree        // guarda arvore atual - pq arvore pode mudar

	return currentNode
}

func (t *PARSER) backtrack(currentNode *NODE, current int, tokens []TOKEN) {
	//backtrack on Tree
	t.tree = currentNode
	t.tree.Edges = nil //reseta os filhos do nó atual para nil

	//backtrack on Token list
	t.lexer.currentPos = current
	t.lexer.currentToken = t.lexer.tokens[t.lexer.currentPos]

}

func (t *PARSER) nextToken() bool {
	t.lexer.currentPos += 1
	t.lexer.currentToken = t.lexer.tokens[t.lexer.currentPos]
	return true
}

func (t *PARSER) term(token string) bool {
	for _, v := range t.lexer.currentToken.classe {
		if v == token {
			t.lexer.currentToken.classe = []string{v}
			node := createNode(token, t.lexer.input[t.lexer.currentPos])
			t.tree.appendNode(node)

			t.nextToken()
			return true
		}
	}
	return false
}
