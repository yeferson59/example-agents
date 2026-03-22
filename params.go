package main

import "github.com/openai/openai-go/v3"

func CreateParams(message string) *openai.ChatCompletionNewParams {
	return &openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(message),
		},
		Tools: []openai.ChatCompletionToolUnionParam{
			openai.ChatCompletionFunctionTool(openai.FunctionDefinitionParam{
				Name:        "get_time",
				Description: openai.String("get the current time"),
				Parameters: openai.FunctionParameters{
					"type": "object",
					"properties": map[string]any{
						"location": map[string]string{
							"type": "string",
						},
					},
					"required": []string{},
				},
			}),
		},
		Seed:  openai.Int(0),
		Model: "x-ai/grok-4.1-fast",
	}
}
