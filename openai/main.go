package openai

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji/v2"
)

type OpenAI struct {
	key   string
	url   string
	model string
}

type Role string

const (
	System    Role = "system"
	Assistant      = "assistant"
	User           = "user"
)

type Message struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

type OpenAIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float32   `json:"temperature"`
}

type ResponseChoice struct {
	Message Message `json:"message"`
}

type OpenAIResponse struct {
	Choices []ResponseChoice `json:"choices"`
}

type OAIReplyContent struct {
	Explanation string `json:"explanation"`
	Code        string `json:"code"`
}

func GetOpenAIClient() OpenAI {
	key := os.Getenv("OPENAI_API_KEY")
	if key == "" {
		color.Red(emoji.Sprint(":cross_mark: OPENAI_API_KEY not set"))
		os.Exit(1)
	}
	return OpenAI{key: key, model: "gpt-3.5-turbo", url: "https://api.openai.com/v1"}
}

func (oai *OpenAI) ChatCompletion(question string) OAIReplyContent {
	// Create OpenAI chat completion request struct
	systemMessage := Message{Role: System, Content: `You are a assistant who is expert in command line tools of linux and unix systems.
Reply in json format with following json structure:
{
	explanation: "explanation will go here",
	code: "code that needs to be execute will go here"
}`}
	userMessage := Message{Role: User, Content: question}
	messages := []Message{systemMessage, userMessage}
	oaiReq := OpenAIRequest{
		Model:       oai.model,
		Messages:    messages,
		Temperature: 0.4,
	}

	// convert to json
	jsonBytes, err := json.Marshal(oaiReq)
	if err != nil {
		log.Fatal(err)
	}

	// Create request object
	req, err := http.NewRequest("POST", oai.url+"/chat/completions", bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}

	// Set content-type
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+oai.key)

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// Read body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonString := string(body)
	var sResp OpenAIResponse
	err = json.Unmarshal([]byte(jsonString), &sResp)
	if err != nil {
		log.Fatal(err)
	}

	var content OAIReplyContent
	contentString := sResp.Choices[0].Message.Content
	err = json.Unmarshal([]byte(contentString), &content)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
