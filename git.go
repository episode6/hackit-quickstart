package main

import (
	"fmt"
	"strings"
)

func assertGitRepo() {
	val, err := execNoPanic("git rev-parse --is-inside-work-tree")
	if err == nil && strings.TrimSpace(val) == "true" {
		return
	}

	shouldGitInit := readConsoleOptionInput("This is not a valid git repo, do you want to 'git init'?", "n", []string{"y", "n"})
	if shouldGitInit == "y" {
		execOrPanic("git init")
	} else {
		panic("hackit-quickstart needs to be called from a valid git repo.")
	}
}

func addGitSubmodule(submodule string, directory string) {
	execOrPanic(fmt.Sprintf("git submodule add \"%v\" ./%v", submodule, directory))
	execOrPanic("git submodule update --init")
}

func readGitOriginURL() string {
	repoURL, err := execNoPanic("git config --get remote.origin.url")
	if err != nil {
		repoURL = readConsolStringInput("Could not find remote 'origin', please enter origin url")
		if repoURL == "" {
			panic("git origin url required")
		}
		execOrPanicWithMessage("git remote add origin "+repoURL, "Failed to add origin to git repo")
		return readGitOriginURL()
	}

	if strings.HasPrefix(repoURL, "git@") {
		repoURL = "https://" + strings.Replace(repoURL[4:len(repoURL)], ":", "/", -1)
	}
	if strings.HasSuffix(repoURL, ".git") {
		repoURL = repoURL[0 : len(repoURL)-4]
	}
	return repoURL
}
