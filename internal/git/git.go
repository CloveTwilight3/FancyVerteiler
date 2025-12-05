package git

import (
	"encoding/json"
	"os"
)

type Service struct {
	cachedCommit  string
	cachedMessage string
}

func New() *Service {
	sha := os.Getenv("GITHUB_SHA")
	if sha == "" {
		sha = "unknown"
	} else {
		sha = sha[:7]
	}

	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	if eventPath == "" {
		return &Service{
			cachedCommit:  sha,
			cachedMessage: "unknown",
		}
	}

	data, err := os.ReadFile(eventPath)
	if err != nil {
		return &Service{
			cachedCommit:  sha,
			cachedMessage: "unknown",
		}
	}

	var event GithubEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return &Service{
			cachedCommit:  sha,
			cachedMessage: "unknown",
		}
	}

	return &Service{
		cachedCommit:  sha,
		cachedMessage: event.HeadCommit.Message,
	}
}

func (s *Service) CommitSHA() string {
	return s.cachedCommit
}

func (s *Service) CommitMessage() string {
	return s.cachedMessage
}
