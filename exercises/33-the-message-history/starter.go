package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

type Message struct {
	Text  string
	Color tcell.Color
	Count int
}

func (m Message) FullText() string {
	if m.Count > 1 {
		return fmt.Sprintf("%s (x%d)", m.Text, m.Count)
	}
	return m.Text
}

type MessageLog struct {
	Messages []Message
}

func (l *MessageLog) AddMessage(text string, color tcell.Color, stack bool) {
	if stack && len(l.Messages) > 0 && l.Messages[len(l.Messages)-1].Text == text {
		l.Messages[len(l.Messages)-1].Count++
		return
	}
	l.Messages = append(l.Messages, Message{Text: text, Color: color, Count: 1})
}

func (l *MessageLog) Render(screen tcell.Screen, x, y, width, height int) {
	renderMessages(screen, x, y, width, height, l.Messages)
}

func renderMessages(screen tcell.Screen, x, y, width, height int, messages []Message) {
	yOffset := height - 1
	for i := len(messages) - 1; i >= 0; i-- {
		msg := messages[i]
		lines := wrap(msg.FullText(), width)
		for j := len(lines) - 1; j >= 0; j-- {
			drawText(screen, x, y+yOffset, lines[j], tcell.StyleDefault.Foreground(msg.Color))
			yOffset--
			if yOffset < 0 {
				return // no more room
			}
		}
	}
}

func wrap(text string, width int) []string {
	var lines []string
	line := ""
	for _, word := range strings.Fields(text) {
		// your turn! a word longer than width must be chopped up first
		switch {
		case line == "":
			line = word
		case len(line)+1+len(word) <= width:
			line += " " + word
		default:
			lines = append(lines, line)
			line = word
		}
	}
	if line != "" {
		lines = append(lines, line)
	}
	return lines
}
