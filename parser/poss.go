package parser

func (t *PARSER) PossP() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("PossP")
	tokens := make([]TOKEN, len(t.lexer.tokens))
	copy(tokens, t.lexer.tokens)
	current := t.lexer.currentPos
	t.lastPos = current

	if !(t.term("POSS") && t.NP()) {

		t.backtrack(currentNode, current, tokens)
		if !(t.NP() && t.term("POSS")) {

			t.backtrack(currentNode, current, tokens)
			if !(t.term("POSS") && t.NumP()) {

				t.backtrack(currentNode, current, tokens)
				return false
			}
		}
	}

	t.keepTrack(lastNode)
	t.tree = lastNode

	return true
}
