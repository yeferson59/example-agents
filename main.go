package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func main() {
	_ = godotenv.Load()
	baseURL := os.Getenv("OPENAI_BASE_URL")
	ctx := context.Background()
	client := openai.NewClient(option.WithBaseURL(baseURL))
	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter some text (press Ctrl+D or Ctrl+Z to end):")

	fmt.Print("Tú: ")
	// Read input line by line
	for scanner.Scan() {
		text := scanner.Text() // Get the current line of text
		if text == "" {
			break // Exit loop if an empty line is entered
		}

		params := CreateParams(text)

		chatCompletion, err := client.Chat.Completions.New(ctx, *params)
		if err != nil {
			panic(err.Error())
		}

		toolCalls := chatCompletion.Choices[0].Message.ToolCalls

		if len(toolCalls) > 0 {
			params.Messages = append(params.Messages, chatCompletion.Choices[0].Message.ToParam())
			for _, toolCall := range toolCalls {
				if toolCall.Function.Name == "get_time" {
					params.Messages = append(params.Messages, openai.ToolMessage(GetTime(), toolCall.ID))
				}
			}

			chatCompletion, err = client.Chat.Completions.New(ctx, *params)
			if err != nil {
				panic(err)
			}

			println("Asistente: ", chatCompletion.Choices[0].Message.Content)
		} else {
			fmt.Print("Tú: ")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}

func GetTime() string {
	return time.Now().String()
}
