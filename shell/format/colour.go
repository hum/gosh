package format

import "fmt"

const (
	ColorReset = 0
	ColorRed   = iota + 30
	ColorGreen
	ColorYellow
	ColorBlue
	ColorPurple
	ColorCyan
	ColorWhite
)

func SetColour(c int, text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", c, text)
}
