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

func (itp *Interpreter) expr() int {
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