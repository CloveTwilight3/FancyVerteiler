package main

import (
	"FancyVerteiler/internal/config"
	"FancyVerteiler/internal/discord"
	"FancyVerteiler/internal/modrinth"

	"github.com/sethvargo/go-githubactions"
)

func main() {
	configPath := githubactions.GetInput("config_path")
	if configPath == "" {
		githubactions.Fatalf("missing input 'config_path'")
	}

	discWebhookURL := githubactions.GetInput("discord_webhook_url")

	githubactions.Infof("Reading config: %s", configPath)

	config.BasePath = "/github/workspace"
	cfg, err := config.ReadFromPath(configPath)
	if err != nil {
		githubactions.Fatalf("failed to read config: %v", err)
	}

	githubactions.Infof("Successfully read config for project: %s", cfg.ProjectName)

	if cfg.Modrinth != nil {
		apiKey := githubactions.GetInput("modrinth_api_key")
		if apiKey == "" {
			githubactions.Fatalf("missing input 'modrinth_api_key'")
		}

		githubactions.Infof("Deploying to Modrinth project: %s", cfg.Modrinth.ProjectID)

		mr := modrinth.New(apiKey)
		if err := mr.Deploy(cfg); err != nil {
			githubactions.Fatalf("failed to deploy to Modrinth: %v", err)
		}
		githubactions.Infof("Successfully deployed to Modrinth project: %s", cfg.Modrinth.ProjectID)
	}

	if discWebhookURL != "" {
		disc := discord.New()
		if err := disc.SendSuccessMessage(discWebhookURL, cfg); err != nil {
			githubactions.Fatalf("failed to send Discord success message: %v", err)
		}
		githubactions.Infof("Successfully sent Discord success message")
	}
}
