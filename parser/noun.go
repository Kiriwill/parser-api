package parser

func (t *PARSER) NP() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("NP")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current

	if !t.Nl() {

		t.backtrack(currentNode, current, tokens)
		return false
	}

	t.keepTrack(lastNode)
	t.tree = lastNode
	return true
}

func (t *PARSER) Nl() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("N'")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current

	if !(t.term("N") && t.PP()) {

		t.backtrack(currentNode, current, tokens)
		if !(t.term("N") && t.CP()) {

			t.backtrack(currentNode, current, tokens)
			if !(t.term("N") && t.Nll()) {

				t.backtrack(currentNode, current, tokens)
				if !(t.AP() && t.Nl()) {

					t.backtrack(currentNode, current, tokens)
					return false
				}
			}
		}
	}

	t.keepTrack(lastNode)
	t.tree = lastNode
	return true
}

func (t *PARSER) Nll() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("N''")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current

	if !(t.CP() && t.Nll()) {

		t.backtrack(currentNode, current, tokens)
		if !(t.PP() && t.Nll()) {

			t.backtrack(currentNode, current, tokens)
			if !(t.AP() && t.Nll()) {

				t.backtrack(currentNode, current, tokens)

				//N'' -> empty
			}
		}
	}

	t.keepTrack(lastNode)
	t.tree = lastNode
	return true
}
