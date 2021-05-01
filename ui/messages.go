package ui

import (
	"fmt"
	"github.com/TwinProduction/go-color"
)

type MessageItem struct {
	Message string
	Purpose MessagePurpose
}

type MessagePurpose int

const (
	None MessagePurpose = iota
	Info
	Warning
	Error
	Success
	ActionRequired
	TroubleshootingTip
)


func DisplayMessage(message MessageItem) {
	fmt.Println(messagePrefix(message.Purpose) + message.Message)
}

func DisplayMessages(messages ...MessageItem) {
	for _, message := range messages {
		DisplayMessage(message)
	}
}

func messagePrefix(purpose MessagePurpose) string {
	switch purpose {
	case 0:
		return ""
	case 1:
		return color.Ize(color.White, "Info: ")
	case 2:
		return color.Ize(color.Yellow, "Warning: ")
	case 3:
		return color.Ize(color.Red, "Error: ")
	case 4:
		return color.Ize(color.Green, "✓ ")
	case 5:
		return color.Ize(color.Yellow, "Action required: ")
	case 6:
		return color.Ize(color.Cyan, "Troubleshooting Tip: ")
	}
	panic(fmt.Sprintf("Unsupported purpose (%v)", purpose))
}
