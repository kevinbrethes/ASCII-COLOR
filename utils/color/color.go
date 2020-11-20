package color

import "strings"

func colorSelected(args []string) string {
	if len(args) < 2 {
		return "default"
	}
	color := strings.ToLower(args[1][8:])
	return color
}

func Choice(args []string) string {
	color := colorSelected(args)

	switch color {
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "yellow":
		return "\033[33m"
	case "blue":
		return "\033[34m"
	case "purple":
		return "\033[35m"
	case "cyan":
		return "\033[36m"
	case "white":
		return "\033[37m"
	default:
		return "\033[0m"
	}
}
