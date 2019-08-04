package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
	"gopy/vm"
)

func main() {
	var isDebug bool
	var debugText string
	flag.BoolVar(&isDebug, "debug", false, "-debug")
	flag.StringVar(&debugText, "expr", "", "-expr=1+1")
	flag.Parse()

	if !isDebug {
		for true {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Zoe> ")
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
	} else {
		if len(debugText) > 0 {
			fmt.Println(vm.Eval(debugText))
		} else {
			panic("add -expr=some expression")
		}
		
	}
}
