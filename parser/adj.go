package parser

func (t *PARSER) AP() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("AP")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current

	if !t.Al() {
		t.backtrack(currentNode, current, tokens)
		return false
	}
	t.tree = lastNode
	return true
}

func (t *PARSER) Al() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("Al")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current

	if !(t.AdvP() && t.Al()) {
		t.backtrack(currentNode, current, tokens)
		if !(t.term("A") && t.PP()) {
			t.backtrack(currentNode, current, tokens)
			if !(t.term("A") && t.CP()) {
				t.backtrack(currentNode, current, tokens)

				if !(t.term("A") && t.All()) {
					t.backtrack(currentNode, current, tokens)

					return false
				}
			}
		}
	}
	t.tree = lastNode
	return true

}

func (t *PARSER) All() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("All")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current

	if !(t.AdvP() && t.All()) {
		t.backtrack(currentNode, current, tokens)
		if !(t.PP() && t.All()) {
			t.backtrack(currentNode, current, tokens)
			//A''-> empty
		}
	}
	t.tree = lastNode
	return true
}
