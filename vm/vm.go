package vm

type Interpreter struct {
	lexer *Lexer
	token *Token
}

func (itp *Interpreter) init(lexer *Lexer) {
	itp.lexer = lexer
	itp.lexer.CurrentToken = new(Token)
	itp.lexer.CurrentChar = new(rune)
	*(itp.lexer).CurrentChar = rune(itp.lexer.Text[itp.lexer.Pos])
	// set to the first token after initialization
	itp.lexer.NextToken()
	itp.token = itp.lexer.CurrentToken
}


func (itp *Interpreter) eat(typ string) {
	if itp.token.Type == typ {
		itp.lexer.NextToken()
	} else {
		panic("type inconsistent")
	}
}

func Eval(text string) int {
	itp := &Interpreter{}
	itp.init(&Lexer{Text: text})

	return itp.expr()
}


func errorSyntax() {
	panic("invalid syntax")
}