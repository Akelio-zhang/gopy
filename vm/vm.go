package vm

import (
	"strconv"
)

type Interpreter struct {
	parser *Parser
	tree AST
}

// interpreter initilize
func (itp *Interpreter) init() {
	itp.parser.init()
}

func (tree BinOp) visit() int {
	token := tree.node
	left := *tree.left
	right := *tree.right
	if token.Type == PLUS {
		return left.visit() + right.visit()
	}
	if token.Type == MINUS {
		return left.visit() - right.visit()
	}
	if token.Type == MUL {
		return left.visit() * right.visit()
	}
	if token.Type == DIV {
		return left.visit() / right.visit()
	}
	panic("syntax error")
}

func (tree UnaryOp) visit() int {
	token := tree.node
	expr := tree.expr
	if token.Type == PLUS {
		return expr.visit()
	} else if token.Type == MINUS {
		return -expr.visit()
	}
	panic("syntax error")
}

func (tree Num) visit() int {
	token := tree.node
	num, err := strconv.Atoi(token.Value)
	if err != nil {
		panic("invalid number")
	}
	return num
}

func (itp *Interpreter) eval() int {
	return itp.tree.visit()
}

func Eval(text string) int {
	lexer := &Lexer{Text: text}
	ps := &Parser{lexer: lexer}
	itp := &Interpreter{parser: ps}
	itp.init()
	tree := ps.parse()
	itp.tree = tree
	return itp.eval()
}

func errorSyntax() {
	panic("invalid syntax")
}