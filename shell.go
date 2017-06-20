package main

import (
	"fmt"
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
	val, err := exec.Command("bash", "-c", command).CombinedOutput()
	if err == nil {
		return string(val)
	}

	if errMessage != "" {
		errMessage = fmt.Sprintf("%v\n", errMessage)
	}
	errMessage = fmt.Sprintf(
		"%vError executing bash command\ncommand: %v\nerror: %v\noutput: %v",
		errMessage,
		command,
		err,
		string(val))
	panic(errMessage)
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
