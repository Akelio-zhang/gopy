package main

import (
	"fmt"
	"bufio"
	"os"
	"gopy/vm"
)

func main() {
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Calc> ")
		text, err := reader.ReadString('\n')
		text = text[0:len(text)-1]
		if err != nil {
			panic("input error!")
		} else if text == "quit" {
			return
		} else {
			fmt.Println(vm.Eval(text))
		}
	}
}
