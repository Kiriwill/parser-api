package parser

func (t *PARSER) PP() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("PP")
	tokens := make([]TOKEN, len(t.lexer.tokens))
	copy(tokens, t.lexer.tokens)
	current := t.lexer.currentPos
	t.lastPos = current

	if !t.PPl() {

		t.backtrack(currentNode, current, tokens)
		return false
	}

	t.keepTrack(lastNode)
	t.tree = lastNode

	return true
}

func (t *PARSER) PPl() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("P'")
	tokens := make([]TOKEN, len(t.lexer.tokens))
	copy(tokens, t.lexer.tokens)
	current := t.lexer.currentPos
	t.lastPos = current

	if !(t.AdvP() && t.PPl()) {

		t.backtrack(currentNode, current, tokens)
		if !(t.term("PREP") && t.DP()) {

			t.backtrack(currentNode, current, tokens)
			if !(t.term("PREP") && t.AdvP()) {

				t.backtrack(currentNode, current, tokens)
				if !(t.term("PREP") && t.CP()) {

					t.backtrack(currentNode, current, tokens)
					if !(t.term("PREP") && t.PP()) { //mudei aqui em 18/11/2021 para priorizar os artigos antes das PREP (pra, na, crase, etc)

						t.backtrack(currentNode, current, tokens)
						if !(t.term("PREP")) {

							t.backtrack(currentNode, current, tokens)
							return false
						}
					}
				}
			}
		}
	}

	t.keepTrack(lastNode)
	t.tree = lastNode

	return true
}
