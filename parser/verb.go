package parser

func (t *PARSER) VP() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("VP")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current
	if !t.Vl() {
		t.backtrack(currentNode, current, tokens)
		return false
	}
	t.tree = lastNode

	return true
}

func (t *PARSER) Vl() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("Vl")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current
	if !(t.AdvP() && t.Vl()) {
		t.backtrack(currentNode, current, tokens)

		if !(t.term("V") && t.DP()) {
			t.backtrack(currentNode, current, tokens)

			if !(t.term("V") && t.PP()) {
				t.backtrack(currentNode, current, tokens)

				if !(t.term("V") && t.CP()) {
					t.backtrack(currentNode, current, tokens)

					if !(t.term("V") && t.AdvP()) {
						t.backtrack(currentNode, current, tokens)

						if !(t.term("V") && t.AP()) {
							t.backtrack(currentNode, current, tokens)

							if !(t.term("V") && t.IP()) {
								t.backtrack(currentNode, current, tokens)

								if !(t.term("V") && t.Vll()) {
									t.backtrack(currentNode, current, tokens)

									return false
								}
							}
						}
					}
				}
			}
		}
	}
	t.tree = lastNode

	return true
}

func (t *PARSER) Vll() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("Vll")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current
	if !(t.AdvP() && t.Vll()) {
		t.backtrack(currentNode, current, tokens)
		if !(t.PP() && t.Vll()) {
			t.backtrack(currentNode, current, tokens)
			//v'' -> empty
		}
	}
	t.tree = lastNode

	return true
}
