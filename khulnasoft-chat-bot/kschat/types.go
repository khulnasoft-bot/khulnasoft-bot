package kschat

import (
	"github.com/khulnasoft-bot/khulnasoft-bot/khulnasoft-chat-bot/kschat/types/chat1"
)

type Result struct {
	Convs []chat1.ConvSummary `json:"conversations"`
}

type SendResponse struct {
	Result chat1.SendRes `json:"result"`
	Error  *Error        `json:"error,omitempty"`
}
