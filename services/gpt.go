package services

import (
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"log"
)

const (
	maxTokens   = 3000
	temperature = 0
	engine      = gpt3.TextDavinci003Engine
)

var client gpt3.Client

func InitClient(api string) {
	client = gpt3.NewClient(api)
}

func GetAnswer(question string) {
	//fmt.Println(apiKey)
	fmt.Print("Answer: ")
	i := 0
	ctx := context.Background()
	if err := client.CompletionStreamWithEngine(ctx, engine, gpt3.CompletionRequest{
		Prompt: []string{
			question,
		},
		MaxTokens:   gpt3.IntPtr(maxTokens),
		Temperature: gpt3.Float32Ptr(temperature),
	}, func(resp *gpt3.CompletionResponse) {
		if i > 1 {
			fmt.Print(resp.Choices[0].Text)
		}
		i++
	}); err != nil {
		log.Fatalln(err)
	}
	fmt.Println()
	fmt.Println()
}

func FormatQuestion(question string) string {
	return "Answer:" + question
}
