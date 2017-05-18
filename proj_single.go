package main

type singleProject struct {
	rootProject
}

func (sp *singleProject) templateAlias() string {
	return "proj/single"
}

func (sp *singleProject) describe() string {
	return "A single-module project"
}

func (sp *singleProject) generate(data *ProjectData) {
	data.Lang.generateLangSpecificFiles(data, "")
	sp.rootProject.generate(data)
	execGdmcResolve()
}
