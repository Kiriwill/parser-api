package parser

func (t *PARSER) QP() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("PPl")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current

	if !(t.term("Q") && t.DP()) {
		t.backtrack(currentNode, current, tokens)
		if !(t.term("Q") && t.PP()) {
			t.backtrack(currentNode, current, tokens)
			if !(t.DP() && t.term("Q")) {
				t.backtrack(currentNode, current, tokens)
				return false
			}
		}
	}
	t.tree = lastNode
	return true
}
