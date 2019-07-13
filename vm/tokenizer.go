package vm

import (
	"unicode"
)

type Token struct {
	Type string
	Value string
}

// tokenizer(scanner) function for lexical analysis 
func (itp *Interpreter) NextToken() {
	// cahr = nil indicates EOF
	for itp.CurrentChar != nil {
		char := *itp.CurrentChar
		if unicode.IsSpace(rune(char)) {
			itp.skipSpace()
			continue
		}

		if unicode.IsDigit(rune(char)) {
			*itp.CurrentToken = itp.getIntegerToken()
			return
		}
	
		if char == '+' {
			itp.advance()
			*itp.CurrentToken = Token{PLUS, string(char)}
			return
		}

		if char == '-' {
			itp.advance()
			*itp.CurrentToken = Token{MINUS, string(char)}
			return 
		}

		panic("parse input error")
	}
	itp.CurrentToken = &Token{EOF, ""}
}

func (itp *Interpreter) getIntegerToken() Token {
	result := ""
	for itp.CurrentChar != nil && unicode.IsDigit(*itp.CurrentChar) {
		result += string(*itp.CurrentChar)
		itp.advance()
	}
	if result != "" {
		return Token{INTEGER, result}
	} else {
		panic("获取整形失败")
	}
	
}


func (itp *Interpreter) skipSpace() {
	for unicode.IsSpace(*itp.CurrentChar) {
		itp.advance()
	}
}

func (itp *Interpreter) advance() {
	itp.Pos += 1
	text := itp.Text
	if itp.Pos > len(text) - 1 {
		itp.CurrentChar = nil
		return
	}
	*itp.CurrentChar = rune(text[itp.Pos])
}