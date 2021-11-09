package parser

func (t *PARSER) NumP() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("NumP")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current

	if !(t.term("NUM") && t.NP()) {

		t.backtrack(currentNode, current, tokens)
		if !(t.term("NUM") && t.PP()) {

			t.backtrack(currentNode, current, tokens)
			return false
		}
	}

	t.keepTrack(lastNode)
	t.tree = lastNode

	return true
}
