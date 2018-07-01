package main

import (
	"path/filepath"
	"testing"
)

func TestGitRepoUrl(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping git-repo-url test in short mode")
	}

	startingDir := getAndPrepWorkingDir(getwd())
	testDir := filepath.Join(startingDir, "git-repo-test")
	prepAndChToProjectTestDir(testDir, true)
	execOrFail("git remote add origin git@github.com:episode6/hackit-gdmc.git", "TestGitRepoUrl", t)

	expected := "https://github.com/episode6/hackit-gdmc"
	result := readGitOriginURL()
	if result != expected {
		t.Errorf("Failed to read correct origin url from git\nExpected: %v\nRead From Git: %v\n", expected, result)
	}
}
