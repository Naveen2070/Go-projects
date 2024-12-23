package handler

import "strings"

// MessageHandler implements the Handler interface.
type MessageHandler struct{}

// ProcessMessage generates a response based on the input message.
func (h *MessageHandler) ProcessMessage(input string) string {
	switch strings.ToLower(input) {
	case "hello":
		return "Hi there! How can I help you?"
	case "bye":
		return "Goodbye! Have a nice day!"
	default:
		return "I don't understand your message."
	}
}
