package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v52/github"
	"golang.org/x/oauth2"
)

func main() {
	githubToken := os.Getenv("DIUN_GH_TOKEN")
	if len(githubToken) == 0 {
		fmt.Println("Github token must be exported as DIUN_GH_TOKEN env var")
		os.Exit(1)
	}

	githubAssignee := os.Getenv("DIUN_GH_ASSIGNEE")
	if len(githubAssignee) == 0 {
		fmt.Println("Github assignee must be exported as DIUN_GH_ASSIGNEE env var")
		os.Exit(1)
	}

	repoOwner := os.Getenv("DIUN_REPO_OWNER")
	if len(repoOwner) == 0 {
		fmt.Println("Github repo owner must be exported as DIUN_REPO_OWNER env var")
		os.Exit(1)
	}

	repoName := os.Getenv("DIUN_REPO_NAME")
	if len(repoName) == 0 {
		fmt.Println("Github repo name be exported as DIUN_REPO_NAME env var")
		os.Exit(1)
	}

	dockerHost := os.Getenv("DIUN_HOSTNAME")
	dockerImage := os.Getenv("DIUN_ENTRY_IMAGE")
	notifTime := os.Getenv("DIUN_ENTRY_CREATED")

	issueTitle := fmt.Sprintf("new version for %s on %s", dockerImage, dockerHost)
	issueBody := fmt.Sprintf(`There is a new image available:

host: %s
image: %s
time registered: %s`, dockerHost, dockerImage, notifTime)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	issueRequest := github.IssueRequest{
		Title:    &issueTitle,
		Body:     &issueBody,
		Assignee: &githubAssignee,
	}

	_, _, err := client.Issues.Create(ctx, repoOwner, repoName, &issueRequest)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("New issue created for %s", dockerImage)
	}
}
