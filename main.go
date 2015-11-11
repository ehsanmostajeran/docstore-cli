package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tsileo/docstore-client"
	"gopkg.in/readline.v1"
)

// FIXME(ts) add history support

func main() {
	dstore := docstore.New("")
	cols, err := dstore.Collections()
	if err != nil {
		panic(err)
	}
	cpls := []*readline.PrefixCompleter{}
	for _, col := range cols {
		cpls = append(cpls, readline.PcItem(col))
	}
	rl, err := readline.NewEx(&readline.Config{
		Prompt: "docstore > ",
		AutoComplete: readline.NewPrefixCompleter(
			readline.PcItem("collection", cpls...),
			// readline.PcItem("say",
			// 	readline.PcItem("hello"),
			// 	readline.PcItem("bye"),
			// ),
			readline.PcItem("query"),
			readline.PcItem("get"),
			readline.PcItem("help"),
		),
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	scol := ""
	var col *docstore.Collection
	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF
			break
		}
		switch {
		case strings.HasPrefix(line, "collection"):
			args := strings.Split(line, " ")
			// A col is specified
			if len(args) == 2 {
				scol = args[1]
				col = dstore.Col(scol)
			}
			println(scol)
		case strings.HasPrefix(line, "get"):
			args := strings.Split(line, " ")
			if len(args) < 2 {
				println("no ID provided")
			}
			id := args[1]
			res := map[string]interface{}{}
			if err := col.Get(id, &res); err != nil {
				fmt.Printf("error: %v\n", err)
			}
			js, _ := json.MarshalIndent(&res, "", "    ")
			fmt.Printf("%s\n", js)
		case strings.HasPrefix(line, "insert"):
			payload := strings.Replace(line, "insert ", "", 1)
			doc := map[string]interface{}{}
			if err := json.Unmarshal([]byte(payload), &doc); err != nil {
				fmt.Printf("error: %s\n", err)
			}
			if err := col.Insert(&doc, nil); err != nil {
				fmt.Printf("error: %s\n", err)
			}
			js, _ := json.MarshalIndent(&doc, "", "    ")
			fmt.Printf("%s\n", js)
		case strings.HasPrefix(line, "query"):
			payload := strings.Replace(line, "query", "", 1)
			q := map[string]interface{}{}
			if err := json.Unmarshal([]byte(payload), &q); err != nil {
				panic(err)
			}
			data, err := col.Iter(q)
			if err != nil {
				fmt.Printf("error: %s\n", err)
			}
			js, _ := json.MarshalIndent(&data, "", "    ")
			fmt.Printf("%s\n", js)
		default:
			println(line)
		}
		// println(line)
	}
}
