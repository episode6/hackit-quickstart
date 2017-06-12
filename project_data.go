package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/episode6/hackit-quickstart/gdmcutil"
	"github.com/episode6/hackit-quickstart/mavenutil"
)

type describable interface {
	describe() string
}

type validator interface {
	validate(data *ProjectData)
}

func validateIfValidator(i interface{}, data *ProjectData) {
	switch v := i.(type) {
	case validator:
		v.validate(data)
	}
}

type projectTemplate interface {
	templatableConfig
	describable
	generate(data *ProjectData)
}

type languageTemplate interface {
	templatableConfig
	describable
	buildscriptDependencies() []string
	generateLangSpecificFiles(data *ProjectData, subdir string)

	GradlePlugins() []string
	JenkinsCommands() []string
}

type packageName string

func (pkg *packageName) asPath() string {
	return filepath.Join(strings.Split(string(*pkg), ".")...)
}

type dependencyResolver interface {
	FormatKeys(keys []string) []string
}

// ProjectData is the object used for our templates
type ProjectData struct {
	Proj        projectTemplate
	Lang        languageTemplate
	Group       packageName
	Version     string
	Name        string
	LicenseName string

	AndroidCompileSdkVersion string
	AndroidBuildToolsVersion string

	gdmcRepoURL string
	depResolver dependencyResolver

	gitRepoURL string
}

func (data *ProjectData) validate() {
	requireSpecial("type", func() bool {
		return data.Proj != nil
	}, func(input string) {
		data.Proj = projectTypes[input]
	})

	requireSpecial("lang", func() bool {
		return data.Lang != nil
	}, func(input string) {
		data.Lang = projectLangs[input]
	})

	requireSpecial("group", func() bool {
		return string(data.Group) != ""
	}, func(input string) {
		data.Group = packageName(input)
	})

	require(&data.Name, "name")
	validateIfValidator(data.Proj, data)
	validateIfValidator(data.Lang, data)
}

func (data *ProjectData) generate() {
	data.Proj.generate(data)
}

// ProjSpecRootGradleBody is used for templating
func (data *ProjectData) ProjSpecRootGradleBody() string {
	return templateTemplateable("root-build.gradle", data.Proj, data)
}

// ProjGradleBody is used for templating
func (data *ProjectData) ProjGradleBody() string {
	return templateAsset("proj-build.gradle", data)
}

// LangSpecProjGradleBody is used for templating
func (data *ProjectData) LangSpecProjGradleBody() string {
	return templateTemplateable("proj-build.gradle", data.Lang, data)
}

// BuildScriptDeps is used for templating
func (data *ProjectData) BuildScriptDeps() []string {
	return data.getDepResolver().FormatKeys(data.Lang.buildscriptDependencies())
}

// DeployableGradleProperties is used for templating
func (data *ProjectData) DeployableGradleProperties() string {
	return templateAsset("deployable-gradle.properties", data)
}

// GitRepoURL is used for templating
func (data *ProjectData) GitRepoURL() string {
	if data.gitRepoURL == "" {
		data.gitRepoURL = readGitOriginURL()
	}
	return data.gitRepoURL
}

// CamelName is used for templating
func (data *ProjectData) CamelName() string {
	return camelCase(data.Name)
}

// CamelNameWithoutApp is used for templating
func (data *ProjectData) CamelNameWithoutApp() string {
	camelName := data.CamelName()
	if strings.HasSuffix(camelName, "Application") {
		return camelName[0 : len(camelName)-11]
	}
	if strings.HasSuffix(camelName, "App") {
		return camelName[0 : len(camelName)-3]
	}
	return camelName
}

func (data *ProjectData) getDepResolver() dependencyResolver {
	if data.depResolver == nil {
		filename := filepath.Join("gdmc", "gdmc.json")
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			// gdmc file doesnt exist, use fallback
			data.depResolver = &mavenutil.MavenResolver{}
		} else {
			data.depResolver = gdmcutil.LoadMap(filename)
		}
	}
	return data.depResolver
}
