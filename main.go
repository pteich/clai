package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/pteich/configstruct"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"

	"github.com/pteich/clai/ai"
	"github.com/pteich/clai/config"
)

func main() {
	var aiRes ai.Response
	var err error

	cfg := config.Config{
		Endpoint: "https://api.groq.com/openai/v1",
		Model:    "llama3-70b-8192",
	}

	opts := make([]configstruct.Option, 0)

	configPath, err := config.FindConfigFile()
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

	err = putils.RunWithDefaultSpinner("🤖 Thinking...", func(spinner *pterm.SpinnerPrinter) error {
		aiRes, err = aiClient.Ask(ctx, phrase)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		pterm.Println(err)
		os.Exit(1)
	}

	pterm.Println(pterm.Magenta(aiRes.Explanation))
	pterm.Println()
	pterm.Println(pterm.Green("$ ") + aiRes.Command)
}
