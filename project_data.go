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

type deployableConfig interface {
	deployableBuildscriptDependencies() []string
	deployableGradlePlugins() []string
	deployableJenkinsCommands() []string
}

type projectTemplate interface {
	templatableConfig
	describable
	generate(data *ProjectData)
}

type languageTemplate interface {
	templatableConfig
	describable

	deployableConfig() deployableConfig

	buildscriptRepos() []string
	projectRepos() []string
	buildscriptDependencies() []string
	generateExtraRootProjectFiles(data *ProjectData)
	generateLangSpecificFiles(data *ProjectData, subdir string)

	GradlePlugins() []string
	JenkinsCommands() []string
}

type packageName string

func (pkg *packageName) asPath() string {
	return filepath.Join(strings.Split(string(*pkg), ".")...)
}

type dependencyResolver interface {
	FindVersion(key string) string
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

	GradleVersion            string
	AndroidSdkDir            string
	AndroidNdkDir            string
	AndroidCompileSdkVersion string

	deployable bool

	gdmcRepoURL string
	depResolver dependencyResolver

	gitRepoURL string
}

// IsDeployable returns true if the project is deployable
func (data *ProjectData) IsDeployable() bool {
	return data.deployable && data.Lang.deployableConfig() != nil
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

// BuildscriptRepos is used for templating
func (data *ProjectData) BuildscriptRepos() []string {
	return append([]string{"jcenter()"}, data.Lang.buildscriptRepos()...)
}

// ProjectRepos is used for templating
func (data *ProjectData) ProjectRepos() []string {
	return append([]string{"jcenter()"}, data.Lang.projectRepos()...)
}

// BuildScriptDeps is used for templating
func (data *ProjectData) BuildScriptDeps() []string {
	deps := data.Lang.buildscriptDependencies()
	if data.IsDeployable() {
		deps = append(deps, "com.episode6.hackit.deployable:deployable")
		deps = append(deps, data.Lang.deployableConfig().deployableBuildscriptDependencies()...)
	}
	deps = append(deps, "com.episode6.hackit.gdmc:gdmc")
	return data.getDepResolver().FormatKeys(deps)
}

// GradlePlugins is used for templating
func (data *ProjectData) GradlePlugins() []string {
	plugins := data.Lang.GradlePlugins()
	if data.IsDeployable() {
		plugins = append(plugins, data.Lang.deployableConfig().deployableGradlePlugins()...)
	}
	plugins = append(plugins, "com.episode6.hackit.gdmc")
	return plugins
}

// DeployableGradleProperties is used for templating
func (data *ProjectData) DeployableGradleProperties() string {
	if data.IsDeployable() {
		return templateAsset("deployable-gradle.properties", data)
	}
	return ""
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

// LookupVersion is used for templating
func (data *ProjectData) LookupVersion(key string) string {
	return data.getDepResolver().FindVersion(key)
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
