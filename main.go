package main

import (
	"gopkg.in/readline.v1"
)

// FIXME(ts) add history support

var completer = readline.NewPrefixCompleter(
	readline.PcItem("say",
		readline.PcItem("hello"),
		readline.PcItem("bye"),
	),
	readline.PcItem("help"),
)

func main() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:       "docstore > ",
		AutoComplete: completer,
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF
			break
		}
		println(line)
	}
}
