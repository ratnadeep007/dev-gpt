package main

import (
	"fmt"
	"os"
	"ratnadeep007/dev-gpt/executor"
	"ratnadeep007/dev-gpt/openai"
	"ratnadeep007/dev-gpt/ui"
	"strings"
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
	// reply := openai.OAIReplyContent{
	// 	Explanation: "some explanation",
	// 	Code:        "ls -alh",
	// }
	s.Stop()

	ui.RenderTextInBox(reply.Code)

	fmt.Println()
	ch := executor.InputPrompt(color.GreenString("Explain or Execute?[e=Explation/x=Execute]:"))

	if strings.ToLower(ch) == "x" {
		executor.Execute(reply.Code)
	} else {
		ui.RenderTextInBox(reply.Explanation)
	}
}
