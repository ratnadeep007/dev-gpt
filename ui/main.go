package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func RenderTextInBox(text string) {
	textStyle := lipgloss.NewStyle().
		PaddingTop(1).
		PaddingBottom(1)

	textStyled := textStyle.Render(text)

	style := lipgloss.NewStyle().
		Bold(true).
		BorderStyle(lipgloss.RoundedBorder()).
		Foreground(lipgloss.Color("#FAFAFA")).
		BorderForeground(lipgloss.Color("63")).
		PaddingTop(1).
		PaddingBottom(1).
		PaddingLeft(1).
		PaddingRight(1).
		Width(lipgloss.Width(textStyled) + 5).
		Align(lipgloss.Center)

	if len(text) > 180 {
		style = style.Width(100)
	}

	fmt.Println(style.Render(text))
}

func InputPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}
