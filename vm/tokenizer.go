package vm

import (
	"unicode"
)

const (
	INTEGER = "INTEGER"
	PLUS = "PLUS"
	MINUS = "MINUS"
	MUL = "MUL"
	DIV = "DIV"
	EOF = "EOF"
)

type Token struct {
	Type string
	Value string
}

type Lexer struct {
	Text string
	Pos int
	CurrentToken *Token
	CurrentChar *rune
}

// tokenizer(scanner) function for lexical analysis 
func (lexer *Lexer) NextToken() {
	// cahr = nil indicates EOF
	for lexer.CurrentChar != nil {
		char := *lexer.CurrentChar
		if unicode.IsSpace(rune(char)) {
			lexer.skipSpace()
			continue
		}

		if unicode.IsDigit(rune(char)) {
			*lexer.CurrentToken = lexer.getIntegerToken()
			return
		}
	
		if char == '+' {
			lexer.advance()
			*lexer.CurrentToken = Token{PLUS, string(char)}
			return
		}

		if char == '-' {
			lexer.advance()
			*lexer.CurrentToken = Token{MINUS, string(char)}
			return 
		}

		if char == '*' {
			lexer.advance()
			*lexer.CurrentToken = Token{MUL, string(char)}
			return 
		}

		if char == '/' {
			lexer.advance()
			*lexer.CurrentToken = Token{DIV, string(char)}
			return 
		}

		panic("parse input error")
	}
	lexer.CurrentToken = &Token{EOF, ""}
}

func (lexer *Lexer) getIntegerToken() Token {
	result := ""
	for lexer.CurrentChar != nil && unicode.IsDigit(*lexer.CurrentChar) {
		result += string(*lexer.CurrentChar)
		lexer.advance()
	}
	if result != "" {
		return Token{INTEGER, result}
	} else {
		panic("get integer token failed")
	}
	
}

func (lexer *Lexer) skipSpace() {
	for unicode.IsSpace(*lexer.CurrentChar) {
		lexer.advance()
	}
}

func (lexer *Lexer) advance() {
	lexer.Pos += 1
	text := lexer.Text
	if lexer.Pos > len(text) - 1 {
		lexer.CurrentChar = nil
		return
	}
	*lexer.CurrentChar = rune(text[lexer.Pos])
}