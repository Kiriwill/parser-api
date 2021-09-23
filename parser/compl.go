package parser

func (t *PARSER) IP() bool {
	node := createNode("IP", "")

	current := t.lexer.currentPos
	t.lastPos = current
	tokens := t.lexer.tokens
	t.tree = node

	if !(t.DP() && t.Il()) {
		t.backtrack(node, current, tokens)

		if !t.Il() {
			t.backtrack(node, current, tokens)
			return false
		}
	}
	t.tree = node
	return true
}

func (t *PARSER) Il() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("Il")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current
	if !(t.Ill() && t.VP()) {
		t.backtrack(currentNode, current, tokens)
		return false
	}
	t.tree = lastNode //arvore atual = arvore anterior
	return true
}

func (t *PARSER) Ill() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("Ill")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current

	if !(t.term("I") && t.Ill()) {
		t.backtrack(currentNode, current, tokens)
		if !(t.DP() && t.Ill()) {
			t.backtrack(currentNode, current, tokens)
			// I'' -> null
		}
	}
	t.tree = lastNode //arvore atual = arvore anterior
	return true
}

func (t *PARSER) CP() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("CP")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current

	if !(t.term("C") && t.IP()) {
		t.backtrack(currentNode, current, tokens)
		return false
	}
	t.tree = lastNode //arvore atual = arvore anterior
	return true
}
