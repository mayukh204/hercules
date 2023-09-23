package git_repo

import (
	"errors"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

func GitClone(url string, directory string) {
	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		log.Printf("Error cloning repository: %v", err)
	}

	// Remove the .git directory
	gitDir := filepath.Join(directory, ".git")
	err = os.RemoveAll(gitDir)
	if err != nil {
		log.Printf("Error removing .git directory: %v", err)
	}
}

func GetRepoNameFromUrl(repoUrl string) (string, error) {
	// Parse the URL using net/url package
	parsedUrl, err := url.Parse(repoUrl)
	if err != nil {
		return "", err
	}

	// Get the path part of the URL (ignoring scheme, www, etc.)
	path := parsedUrl.Path

	// Remove the leading "/" from the path
	trimmedPath := strings.TrimPrefix(path, "/")

	// Split the path into parts
	parts := strings.Split(trimmedPath, "/")

	// Validate that the URL has at least two parts (username and repo)
	if len(parts) < 2 {
		return "", errors.New("invalid GitHub URL")
	}

	// Get the username and repository name
	username := parts[0]
	repoName := parts[1]

	// Combine them to get the repository directory
	repoDir := username + "/" + repoName

	return repoDir, nil
}
