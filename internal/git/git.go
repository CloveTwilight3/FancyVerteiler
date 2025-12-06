package git

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

type Service struct {
	cachedCommit  string
	cachedMessage string
}

func New() *Service {
	sha, err := fetchLatestCommitSHA()
	if err != nil {
		sha = "unknown"
	}

	msg, err := fetchLatestCommitMessage()
	if err != nil {
		msg = "unknown"
	}

	return &Service{
		cachedCommit:  sha,
		cachedMessage: msg,
	}
}

func (s *Service) CommitSHA() string {
	return s.cachedCommit
}

func (s *Service) CommitMessage() string {
	return s.cachedMessage
}

func fetchLatestCommitSHA() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--short", "HEAD")
	cmd.Dir = os.Getenv("GITHUB_WORKSPACE")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return out.String(), nil
}

func fetchLatestCommitMessage() (string, error) {
	cmd := exec.Command("git", "log", "-1", "--pretty=%B")
	cmd.Dir = os.Getenv("GITHUB_WORKSPACE")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return strings.TrimSpace(out.String()), nil
}
