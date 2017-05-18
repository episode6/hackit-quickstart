package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func execGdmcResolve() {
	execOrPanic("./gradlew -Pgdmc.forceResolve=true gdmcResolve")
}

func execOrPanic(command string) string {
	return execOrPanicWithMessage(command, "")
}

func execOrPanicWithMessage(command string, errMessage string) string {
	val, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		if errMessage == "" {
			panic(err)
		}
		panic(errMessage)
	}
	return string(val)
}

func mkdir(paths ...string) {
	for _, path := range paths {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func pathWithOptSubdir(subdir string, pathSegments ...string) string {
	fullpath := filepath.Join(pathSegments...)
	if subdir == "" {
		return fullpath
	}
	return filepath.Join(subdir, fullpath)
}
