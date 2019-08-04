package vm

type Parser struct {
	lexer *Lexer
	token *Token
}

type AST interface {
	visit() int
 }

type Tree struct {
	left *Tree
	node Token
	tp string
	right *Tree
}

func (ps *Parser) init() {
	ps.lexer.CurrentToken = new(Token)
	ps.lexer.CurrentChar = new(rune)
	*(ps.lexer).CurrentChar = rune(ps.lexer.Text[ps.lexer.Pos])
	// set to the first token after initialization
	ps.lexer.NextToken()
	ps.token = ps.lexer.CurrentToken
}

/* Parser function is generated according to context-free grammar(EBNF)
   and the transformation guidelines.
*/
func (ps *Parser) eat(typ string) {
	if ps.token.Type == typ {
		ps.lexer.NextToken()
	} else {
		panic("type inconsistent")
	}
}

func (ps *Parser) factor() Tree {
	token := *ps.token
	if token.Type == INTEGER {
		ps.eat(INTEGER)
		return Tree{node: token, tp: NUM}
	} else if token.Type == LPAREN {
		ps.eat(LPAREN)
		exprTree := ps.expr()
		ps.eat(RPAREN)
		return exprTree
	}
	panic("invalid syntax")
}

func (ps *Parser) term() Tree {
	tree := ps.factor()
	for ps.token.Type == MUL || ps.token.Type == DIV {
		token := *ps.token
		if token.Type == MUL {
			ps.eat(MUL)
		} else if token.Type == DIV {
			ps.eat(DIV)
		}
		left := tree
		right := ps.factor()
		tree = Tree{left: &left, node: token, tp: BINOP, right: &right}
	}
	return tree
}

func (ps *Parser) expr() Tree {
	tree := ps.term()
	for ps.token.Type == PLUS || ps.token.Type == MINUS {
		token := *ps.token
		if token.Type == PLUS {
			ps.eat(PLUS)
		} else if token.Type == MINUS {
			ps.eat(MINUS)
		}
		left := tree
		right := ps.term()
		tree = Tree{left: &left, node: token, tp: BINOP, right: &right}
	}
	
	return tree
}

func(ps *Parser) parse() Tree {
	return ps.expr()
}