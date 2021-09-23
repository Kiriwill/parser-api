package parser

func (t *PARSER) AdvP() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("AdvP")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current

	if !t.Advl() {
		t.backtrack(currentNode, current, tokens)
		return false
	}
	t.tree = lastNode
	return true
}

func (t *PARSER) Advl() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("Advl")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current

	if !(t.term("ADV") && t.PP()) {
		t.backtrack(currentNode, current, tokens)
		if !(t.term("ADV") && t.Advll()) {
			t.backtrack(currentNode, current, tokens)
			return false
		}
	}
	t.tree = lastNode
	return true
}

func (t *PARSER) Advll() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("Advll")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current

	if !(t.Advl() && t.Advll()) {
		t.backtrack(currentNode, current, tokens)
		//ADV' -> null
	}
	t.tree = lastNode
	return true
}
