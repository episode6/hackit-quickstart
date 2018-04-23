//go:generate go-bindata templates/...
package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/namsral/flag"
)

// AppVersion is the current version of this app
const AppVersion = "0.0.6"

var projectTypes = map[string]projectTemplate{
	"single": &singleProject{},
	"multi":  &multiProject{},
	"sub":    &subProject{},
}

var projectLangs = map[string]languageTemplate{
	"java":             &javaLibrary{},
	"groovy":           &groovyLibrary{},
	"gradle":           &gradleLibrary{},
	"android":          &androidLibrary{},
	"androidApp":       &androidApplication{},
	"androidAppDagger": &androidApplicationWithDagger{},
	// "androidAppBootstrap": &androidApplicationWithBootstrap{},
}

func main() {
	flag.String(
		flag.DefaultConfigFlagname, defaultConfigFilePath(),
		"path to config file")

	versionFlag := flag.Bool("v", false, "Display hackit-quickstart version")

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
		"version", defaultProjectVersion,
		"Initial version name to use")
	nameStr := flag.String(
		"name", "",
		"The name of the new module to generate (for a multi-module project, this will be the sub-module's name)")
	gdmcRepoURLStr := flag.String(
		"gdmc", "",
		"Url of a shared gdmc repo to add as a sub-module")
	deployable := flag.Bool(
		"deployable", false,
		"Make a deployable library (has no effect on apps)")
	licenseNameStr := flag.String(
		"licenseName", defaultLicenseName,
		"The name of the license you want to use (for deployable libraries)")
	gradleVersion := flag.String(
		"gradleVersion", defaultGradleVersion,
		"Gradle version to apply to the project (root project only)")
	androidSdkDirStr := flag.String(
		"androidSdkDir", defaultAndroidSdkDir(),
		"Android sdk directory")
	androidNdkDirStr := flag.String(
		"androidNdkDir", defaultAndroidNdkDir(),
		"Android ndk directory")
	androidCompileSdkVersionStr := flag.String(
		"androidCompileSdkVersion", defaultAndroidCompileSdkVersion,
		"For android apps/libs, the value of compileSdkVersion")

	flag.Parse()

	if *versionFlag {
		fmt.Printf("hackit-quickstart v%v\n", AppVersion)
		os.Exit(0)
	}

	data := &ProjectData{
		Proj:        projectTypes[*projectTypeString],
		Lang:        projectLangs[*projectLangString],
		Group:       packageName(*groupStr),
		Version:     *versionStr,
		Name:        *nameStr,
		LicenseName: *licenseNameStr,
		gdmcRepoURL: *gdmcRepoURLStr,
		deployable:  *deployable,

		GradleVersion:            *gradleVersion,
		AndroidSdkDir:            *androidSdkDirStr,
		AndroidNdkDir:            *androidNdkDirStr,
		AndroidCompileSdkVersion: *androidCompileSdkVersionStr,
	}

	performProjectGeneration(data)
}

func performProjectGeneration(data *ProjectData) {
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

func readConsolStringInput(prompt string) string {
	fmt.Printf("%v\n: ", prompt)
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(input)
}

func readConsoleOptionInput(prompt string, defaultOption string, options []string) string {
	var fullPrompt = prompt + "\n("
	for i, opt := range options {
		if defaultOption == opt {
			fullPrompt += strings.ToUpper(opt)
		} else {
			fullPrompt += opt
		}
		if i != len(options)-1 {
			fullPrompt += "/"
		}
	}
	fullPrompt += "): "
	fmt.Print(fullPrompt)

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return defaultOption
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return defaultOption
	}

	for _, opt := range options {
		if input == opt {
			return opt
		}
	}
	return readConsoleOptionInput(prompt, defaultOption, options)
}
