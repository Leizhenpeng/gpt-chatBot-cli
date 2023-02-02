package services

import (
	"bufio"
	"fmt"
	"os"
)

func ReadUserInput() string {
	//fmt scan
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return "exit"
}

func ReadUserInputWithPrompt(prompt string) string {
	fmt.Print(prompt)
	return ReadUserInput()
}

func AskUserQuestion() string {
	return ReadUserInputWithPrompt("You:")
}
