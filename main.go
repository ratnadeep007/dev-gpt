package main

import (
	"os"
	"ratnadeep007/dev-gpt/openai"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/kyokomi/emoji/v2"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		color.Red(emoji.Sprint(":cross_mark: Question required"))
		os.Exit(1)
	}

	oai := openai.GetOpenAIClient()

	s := spinner.New(spinner.CharSets[9], 100*time.Microsecond)
	s.Suffix = " Thinking..."
	s.Color("green")
	s.Start()
	reply := oai.ChatCompletion(args[1])
	// fmt.Println(reply)
	s.FinalMSG = reply
	s.Stop()
}
