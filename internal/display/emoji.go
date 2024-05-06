package display

import (
	"fmt"
)

// Emoji turns the text title into an emoticon
func Emoji(emoji string) string {
	switch emoji {
	case "art":
		return "🎨"
	case "guitar":
		return "🎸"
	case "drums":
		return "🥁"
	case "hourglass":
		return "⏳"
	case "microphone":
		return "🎤"
	case "piano":
		return "🎹"
	case "star":
		return "⭐"
	case "speaker":
		return "🔉"
	default:
		return fmt.Sprintf(":%s:", emoji)
	}
}
