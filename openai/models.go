package openai

type OpenAI struct {
	key   string
	url   string
	model string
}

type Role string

const (
	System    Role = "system"
	Assistant Role = "assistant"
	User      Role = "user"
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
