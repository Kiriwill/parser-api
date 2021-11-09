package parser

func (t *PARSER) DP() bool {
	// passar essa esrtutura para as outras.
	// usar os ponteiros para ir caminhando na arvore
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("DP")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current

	if !(t.term("D") && t.NP()) {

		t.backtrack(currentNode, current, tokens)

		if !(t.term("D") && t.PossP()) {

			t.backtrack(currentNode, current, tokens)

			if !(t.term("D") && t.NumP()) {

				t.backtrack(currentNode, current, tokens)

				if !(t.NP()) {

					t.backtrack(currentNode, current, tokens)

					if !(t.PossP()) {

						t.backtrack(currentNode, current, tokens)

						if !(t.NumP()) {

							t.backtrack(currentNode, current, tokens)

							if !(t.term("D")) {

								t.backtrack(currentNode, current, tokens)
								return false
							}
						}

					}

				}

			}
		}
	}
	t.keepTrack(lastNode)
	t.tree = lastNode //volta para endereço do pai
	return true
}

func (t *PARSER) Dl() bool {
	lastNode := t.tree // guarda endereço da arvore anterior

	currentNode := t.nextNode("D'")
	tokens := t.lexer.tokens
	current := t.lexer.currentPos
	t.lastPos = current
	if !(t.term("D")) {

		t.backtrack(currentNode, current, tokens)
		return false
	}

	t.keepTrack(lastNode)
	t.tree = lastNode

	return true
}
