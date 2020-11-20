package errors

import (
	"fmt"
)

func noArguments(args []string) bool { //handling "no arguments" error
	if len(args) == 0 {
		return true
	}
	return false
}

func manyArguments(args []string) bool {
	if len(args) > 2 {
		return true
	}
	return false
}

func invalidCharacter(args []string) bool { //handling "non-valid character" error
	for _, val := range args {
		for i := range val {
			if val[i] < 32 || val[i] > 126 {
				return true
			}
		}
	}
	return false
}

func colorError(args []string) bool { //handling invalid argument
	if len(args) < 2 {
		return true
	}

	if len(args[1]) < 9 {
		return true
	}

	if args[1][0:8] != "--color=" {
		return true
	}
	return false
}

func HandlingError(args []string) bool {
	switch {
	case noArguments(args):
		fmt.Println("You didn't give any arguments !")
		return true
	case invalidCharacter(args):
		fmt.Println("You used a non-valid character.")
		return true
	case manyArguments(args):
		fmt.Println("You have given too many arguments !")
		return true
	case colorError(args):
		fmt.Println("Invalid syntax -> <text> --color=<color>\n")
		fmt.Println("Color list : ")
		fmt.Println("\033[31m", "- red")
		fmt.Println("\033[32m", "- green")
		fmt.Println("\033[33m", "- yellow")
		fmt.Println("\033[34m", "- blue")
		fmt.Println("\033[35m", "- purple")
		fmt.Println("\033[36m", "- cyan")
		fmt.Println("\033[37m", "- white")
		return true
	default:
		return false
	}
}
