package parser

func (t *PARSER) VP() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("VP")
	tokens := make([]TOKEN, len(t.lexer.tokens))
	copy(tokens, t.lexer.tokens)
	current := t.lexer.currentPos
	t.lastPos = current
	if !t.Vl() {

		t.backtrack(currentNode, current, tokens)
		return false
	}

	t.keepTrack(lastNode)
	t.tree = lastNode
	return true
}

func (t *PARSER) Vl() bool {
	// Toda sentença precisa ter um ou mais verbos
	// Se eu passei pelo verbo e não encontrei o final da sentença
	// significa que encontrei um caminho sem volta na árvore
	// nao há como alguem ter o verbo como complemento,
	// isto é, abaixo da arvore se não for uma outra sentença
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("V'")
	tokens := make([]TOKEN, len(t.lexer.tokens))
	copy(tokens, t.lexer.tokens)
	current := t.lexer.currentPos
	t.lastPos = current

	if !(t.AdvP() && t.Vl()) || (t.lastPos != (len(t.lexer.input) - 1)) {

		t.backtrack(currentNode, current, tokens)

		if !(t.term("V") && t.PP()) || t.lastPos != (len(t.lexer.input)-1) {

			t.backtrack(currentNode, current, tokens)

			if !(t.term("V") && t.AdvP()) || t.lastPos != (len(t.lexer.input)-1) {

				t.backtrack(currentNode, current, tokens)

				if !(t.term("V") && t.AP()) || t.lastPos != (len(t.lexer.input)-1) {

					t.backtrack(currentNode, current, tokens)

					if !(t.term("V") && t.CP()) || t.lastPos != (len(t.lexer.input)-1) {

						t.backtrack(currentNode, current, tokens)

						if !(t.term("V") && t.IP()) || t.lastPos != (len(t.lexer.input)-1) {

							t.backtrack(currentNode, current, tokens)

							if !(t.term("V") && t.DP()) || t.lastPos != (len(t.lexer.input)-1) {

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

	t.keepTrack(lastNode)
	t.tree = lastNode

	return true
}

func (t *PARSER) Vll() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("V''")
	tokens := make([]TOKEN, len(t.lexer.tokens))
	copy(tokens, t.lexer.tokens)
	current := t.lexer.currentPos
	t.lastPos = current
	if !(t.AdvP() && t.Vll()) {

		t.backtrack(currentNode, current, tokens)
		if !(t.PP() && t.Vll()) {

			t.backtrack(currentNode, current, tokens)
			//v'' -> empty
		}
	}

	t.keepTrack(lastNode)
	t.tree = lastNode

	return true
}
