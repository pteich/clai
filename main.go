package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/pteich/configstruct"

	"github.com/pteich/clai/ai"
)

func main() {
	cfg := Config{
		Endpoint: "https://api.groq.com/openai/v1",
		Model:    "llama3-70b-8192",
	}

	opts := make([]configstruct.Option, 0)

	configPath, err := FindConfigFile()
	if err == nil {
		opts = append(opts, configstruct.WithYamlConfig(configPath))
	}

	err = configstruct.Parse(&cfg, opts...)
	if err != nil {
		fmt.Println("Failed to parse config:", err)
		fmt.Println("Usage: clai <prompt>")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: clai <prompt>")
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	defer cancel()

	aiClient := ai.New(cfg.Token, cfg.Endpoint, cfg.Model, cfg.Shell, cfg.OS)

	phrase := strings.Join(os.Args[1:], " ")

	aiRes, err := aiClient.Ask(ctx, phrase)
	if err != nil {
		fmt.Println("Failed to get AI response:", err)
		os.Exit(1)
	}

	fmt.Printf("\x1b[1;35m%s\x1b[0m", aiRes.Explanation)
	fmt.Println()
	fmt.Printf("\n\x1b[1;32m$ \x1b[0m%s\n", aiRes.Command)
}
