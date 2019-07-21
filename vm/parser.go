package vm

import (
	"strconv"
)

/* Parser function is generated according to context-free grammar(EBNF)
   and the transformation guidelines.
*/
func (itp *Interpreter) factor() int {
	num, err := strconv.Atoi(itp.token.Value)
	if err != nil {
		panic("invalid number")
	}
	itp.eat(INTEGER)
	return num
}

func (itp *Interpreter) term() int {
	result := itp.factor()
	for itp.token.Type == MUL || itp.token.Type == DIV {
		token := itp.token
		if token.Type == MUL {
			itp.eat(MUL)
			result = result * itp.factor()
		} else if token.Type == DIV {
			itp.eat(DIV)
			result = result / itp.factor()
		}
	}
	return result
}

func (itp *Interpreter) expr() int {
	result := itp.term()
	for itp.token.Type == PLUS || itp.token.Type == MINUS {
		token := itp.token
		if token.Type == PLUS {
			itp.eat(PLUS)
			result = result + itp.term()
		} else if token.Type == MINUS {
			itp.eat(MINUS)
			result = result - itp.term()
		}
	}
	return result
}