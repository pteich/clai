package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type AI struct {
	client *openai.Client
	model  string
	shell  string
	os     string
}

func New(token string, endpoint string, model string, shell string, platform string) *AI {
	clientConfig := openai.DefaultConfig(token)
	if endpoint != "" {
		clientConfig.BaseURL = endpoint
	}

	clientConfig.AssistantVersion = "v2"

	client := openai.NewClientWithConfig(clientConfig)

	if platform != "" {
		platform := runtime.GOOS
		if strings.Contains(platform, "darwin") {
			platform = "macOS"
		}
	}

	if shell == "" {
		shellData := strings.Split(os.Getenv("SHELL"), "/")
		shell = shellData[len(shellData)-1]
	}

	return &AI{
		client: client,
		model:  model,
		os:     platform,
		shell:  shell,
	}
}

func (a *AI) getSystemPrompt() string {
	prompt := fmt.Sprintf(`
		You are CLAI, an AI CLI code assistant that only responds with '%s' shell command line instructions for the '%s' operating system.
		You do not provide any other information or comments. Given a user query, respond with the most
		relevant CLI command to accomplish what the user is asking for and nothing else. Make sure the command is guaranteed to work on the given shell and OS.
		Ignore any pleasantries, commentary or questions from the user and only respond with the '%s' command for '%s' and a
		short one sentence description explaining the command. You can use the internet to find the command and explanation.
		Don't write any code or markdown. If the CLI command consists of multiple lines, separate each line with a newline character.
		Return the data in a JSON format like this { \"command\": \"command_here\", \"explanation\": \"explanation_here\" }
		If you can't find a command, still respond with the same JSON, but leave command empty and explain why you couldn't find anything.
		`,
		a.shell, a.os, a.shell, a.os,
	)

	return prompt
}

func (a *AI) Ask(ctx context.Context, prompt string) (Response, error) {
	var aiResponse Response

	req := openai.ChatCompletionRequest{
		Model: a.model,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: a.getSystemPrompt()},
			{Role: "user", Content: prompt},
		},
	}

	res, err := a.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return aiResponse, fmt.Errorf("failed to get AI response: %w", err)
	}

	if len(res.Choices) == 0 {
		return aiResponse, fmt.Errorf("no response from AI")
	}

	err = json.Unmarshal([]byte(res.Choices[0].Message.Content), &aiResponse)
	if err != nil {
		return aiResponse, fmt.Errorf("failed to parse AI response: '%s' %w", res.Choices[0].Message.Content, err)
	}

	return aiResponse, nil
}
