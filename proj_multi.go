package main

type multiProject struct {
	rootProject
	subProject
}

func (mp *multiProject) templateAlias() string {
	return "proj/multi"
}

func (mp *multiProject) describe() string {
	return "A multi-module project with a single sub-module to start"
}

func (mp *multiProject) validate(data *ProjectData) {
	validateIfValidator(mp.rootProject, data)
	validateIfValidator(mp.subProject, data)
}

func (mp *multiProject) generate(data *ProjectData) {
	mp.rootProject.generate(data)

	mp.subProject.skipAmmendSettings = true
	mp.subProject.generate(data)
}
