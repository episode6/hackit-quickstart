//go:generate go-bindata templates/...
package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

var projectTypes = map[string]projectTemplate{
	"single": &singleProject{},
	"multi":  &multiProject{},
	"sub":    &subProject{},
}

var projectLangs = map[string]languageTemplate{
	"java":       &javaLibrary{},
	"groovy":     &groovyLibrary{},
	"gradle":     &gradleLibrary{},
	"android":    &androidLibrary{},
	"androidApp": &androidApplication{},
}

func main() {
	defaultGdmcRepo := os.Getenv("GDMC")

	projectTypeString := flag.String(
		"type", "",
		fmt.Sprint("Type of project to create. Valid values are ", describeProjectTypes()))
	projectLangString := flag.String(
		"lang", "",
		fmt.Sprint("Language of project to create. Valid values are ", describeProjectLanguages()))
	groupStr := flag.String(
		"group", "",
		"GroupId (aka package name) of library to generate")
	versionStr := flag.String(
		"version", "0.0.1-SNAPSHOT",
		"Initial version name to use")
	nameStr := flag.String(
		"name", "",
		"The name of the new module to generate")
	gdmcRepoURLStr := flag.String(
		"gdmc", defaultGdmcRepo,
		"Url of a shared gdmc repo to add as a sub-module")
	noGdmc := flag.Bool(
		"noGdmcRepo", false,
		"Don't use a gdmc repo, equivilent to gdmc=\"\"")

	flag.Parse()

	if *noGdmc {
		*gdmcRepoURLStr = ""
	}

	data := &ProjectData{
		Proj:        projectTypes[*projectTypeString],
		Lang:        projectLangs[*projectLangString],
		Group:       packageName(*groupStr),
		Version:     *versionStr,
		Name:        *nameStr,
		gdmcRepoURL: *gdmcRepoURLStr,
	}

	performProjectGeneration(data)
}

func performProjectGeneration(data *ProjectData) {
	assertGitRepo()
	data.validate()
	data.generate()
}

func describeProjectTypes() string {
	fullDesc := ""
	for _, v := range reflect.ValueOf(projectTypes).MapKeys() {
		name := fmt.Sprintf("%v", v)
		desc := projectTypes[name].describe()
		fullDesc += "\n\t\t" + name + ": " + desc
	}
	return fullDesc
}

func describeProjectLanguages() string {
	fullDesc := ""
	for _, v := range reflect.ValueOf(projectLangs).MapKeys() {
		name := fmt.Sprintf("%v", v)
		desc := projectLangs[name].describe()
		fullDesc += "\n\t\t" + name + ": " + desc
	}
	return fullDesc
}

func require(a *string, flagName string) {
	if a != nil && *a != "" {
		return
	}
	*a = readMissingParam(flagName)
	require(a, flagName)
}

func requireSpecial(flagName string, check func() bool, set func(input string)) {
	for !check() {
		input := readMissingParam(flagName)
		set(input)
	}
}

func readMissingParam(flagName string) string {
	flg := flag.Lookup(flagName)
	fmt.Printf("Missing or invalid parameter '%v' - %v\nEnter %v: ", flg.Name, flg.Usage, flg.Name)
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(input)
}
