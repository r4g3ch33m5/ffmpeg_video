package adapter

type ChatCompletion interface {
	ChatCompletion(prompt string, model string, maxTokens int) (string, error)
}
