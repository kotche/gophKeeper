package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

func executor(in string) {
	in = strings.TrimSpace(in)
	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "exit":
		fmt.Println("GophKeeper stop")
		os.Exit(0)
		return
	case "reg":
		fmt.Println("i input reg")
	case "auth":
		fmt.Println("i input auth")
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	var s []prompt.Suggest

	switch d.Text {
	case "reg":
		s = []prompt.Suggest{
			{Text: "reg", Description: "user registration"},
		}
	case "auth":
		s = []prompt.Suggest{
			{Text: "auth", Description: "user authentication"},
		}
	case "exit":
		s = []prompt.Suggest{
			{Text: "exit", Description: "exit program"},
		}
	default:
		s = []prompt.Suggest{
			{Text: "reg", Description: "user registration"},
			{Text: "auth", Description: "user authentication"},
			{Text: "exit", Description: "exit program"},
		}
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	fmt.Println("GophKeeper start")

	p := prompt.New(
		executor,
		completer,
		prompt.OptionTitle("menu"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Green),
	)
	p.Run()
}
