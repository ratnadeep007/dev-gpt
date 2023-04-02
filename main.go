package main

import (
	"fmt"
	"log"
	"os"
	"ratnadeep007/dev-gpt/openai"
)

func main() {
  args := os.Args

  if len(args) < 2 {
    log.Fatal("Question is required")
  }

  oai := openai.GetOpenAIClient()
  reply := oai.ChatCompletion(args[1])
  fmt.Println(reply)
}
