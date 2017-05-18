package main

import (
	"fmt"
	"strings"
)

func assertGitRepo() {
	val := execOrPanicWithMessage("git rev-parse --is-inside-work-tree", "Not in a git repo")
	if strings.TrimSpace(val) != "true" {
		panic("Not in a git repo")
	}
}

func addGitSubmodule(submodule string, directory string) {
	execOrPanic(fmt.Sprintf("git submodule add \"%v\" ./%v", submodule, directory))
	execOrPanic("git submodule update --init")
}

func readGitOriginURL() string {
	repoURL := strings.TrimSpace(execOrPanic("git config --get remote.origin.url"))
	if strings.HasPrefix(repoURL, "git@") {
		repoURL = "https://" + strings.Replace(repoURL[4:len(repoURL)], ":", "/", -1)
	}
	if strings.HasSuffix(repoURL, ".git") {
		repoURL = repoURL[0 : len(repoURL)-4]
	}
	return repoURL
}
