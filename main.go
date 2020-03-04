package main

import (
	"AppSync/github"
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) < 2 {
		panic("fatal: required args - githubUserName & appName")
	}

	githubUserName := os.Args[1]
	appName := os.Args[2]
	userInstance := github.Init(githubUserName, appName)
	urls := github.GetMatchingGistsUrls(userInstance)

	for _, url := range urls {
		fmt.Print(url)
	}
}
