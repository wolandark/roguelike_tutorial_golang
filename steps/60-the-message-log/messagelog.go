package main

import "github.com/gdamore/tcell/v2"

type Message struct {
	Text  string
	Color tcell.Color
}

type MessageLog struct {
	Messages []Message
}

func (l *MessageLog) AddMessage(text string, color tcell.Color) {
	l.Messages = append(l.Messages, Message{Text: text, Color: color})
}

func (l *MessageLog) Render(screen tcell.Screen, x, y, width, height int) {
	renderMessages(screen, x, y, width, height, l.Messages)
}

func renderMessages(screen tcell.Screen, x, y, width, height int, messages []Message) {
	yOffset := height - 1
	for i := len(messages) - 1; i >= 0 && yOffset >= 0; i-- {
		msg := messages[i]
		drawText(screen, x, y+yOffset, msg.Text, tcell.StyleDefault.Foreground(msg.Color))
		yOffset--
	}
}
