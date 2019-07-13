package vm

import (
	"strconv"
)

const (
	INTEGER = "INTEGER"
	PLUS = "PLUS"
	MINUS = "MINUS"
	EOF = "EOF"
)

type Interpreter struct {
	Text string
	Pos int
	CurrentToken *Token
	CurrentChar *rune
}

func (itp *Interpreter) init() {
	itp.CurrentToken = new(Token)
	itp.CurrentChar = new(rune)
	*itp.CurrentChar = rune(itp.Text[itp.Pos])
}


func (itp *Interpreter) eat(typ string) {
	if itp.CurrentToken.Type == typ {
		itp.NextToken()
	} else {
		panic("类型不合法！")
	}
}

func Eval(text string) int {
	itp := &Interpreter{Text: text}
	itp.init()
	// here is parsing (syntax analysis)
	itp.NextToken()
	left,_ := strconv.Atoi(itp.CurrentToken.Value)
	itp.eat(INTEGER)

	op := itp.CurrentToken.Value
	if op == "+" {
		itp.eat(PLUS)
	} else {
		itp.eat(MINUS)
	}
	
	right,_ := strconv.Atoi(itp.CurrentToken.Value)
	itp.eat(INTEGER)
	 
	// here is interpreting
	if op == "+" {
		return left + right
	}

	if op == "-" {
		return left - right
	}
	
	panic("不支持的语法")
}