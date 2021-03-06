package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/episode6/hackit-quickstart/mavenutil"
)

const testingRepoURL = "https://git.idevsix.com:ghackett/test"
const hackitGdmcRepo = "git@github.com:episode6/hackit-gdmc.git"

var mavenResolver = &mavenutil.MavenResolver{}

func TestSingleProjectGenerationSimple(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	doTestOnEachLang(t, "single_simple", true, func(testName string, testLang languageTemplate) {
		data := makeDefaultProjectData(&ProjectData{
			Proj:  &singleProject{},
			Lang:  testLang,
			Group: packageName("com.g6init.testing"),
			Name:  "some-test-product",

			gdmcRepoURL: hackitGdmcRepo,
			deployable:  false,
		})

		t.Logf("Generating project: %v", testName)
		performProjectGeneration(data)
		execOrFail("./gradlew clean assemble check", testName, t)
	})
}

func TestSingleProjectGenerationDeployable(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	doTestOnEachLang(t, "single_deployable", true, func(testName string, testLang languageTemplate) {
		if testLang.deployableConfig() == nil {
			t.Logf("Skipping non-deployable test: %v", testName)
			return
		}
		data := makeDefaultProjectData(&ProjectData{
			Proj:  &singleProject{},
			Lang:  testLang,
			Group: packageName("com.g6init.testing"),
			Name:  "some-test-product",

			gdmcRepoURL: hackitGdmcRepo,
			deployable:  true,
		})

		t.Logf("Generating project: %v", testName)
		performProjectGeneration(data)
		execOrFail("./gradlew clean assemble check", testName, t)
	})
}

func TestMultiProjectGenerationSimple(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	doTestOnEachLang(t, "multi_simple", true, func(testName string, testLang languageTemplate) {
		multiProjectData := makeDefaultProjectData(&ProjectData{
			Proj:  &multiProject{},
			Lang:  testLang,
			Group: packageName("com.g6init.testing"),
			Name:  "some-test-product",

			gdmcRepoURL: hackitGdmcRepo,
			deployable:  false,
		})

		t.Logf("Generating root project for project: %v", testName)
		performProjectGeneration(multiProjectData)
		execOrFail("./gradlew clean assemble check", testName, t)

		subProjectData := makeDefaultProjectData(&ProjectData{
			Proj:       &subProject{},
			Lang:       testLang,
			Group:      packageName("com.g6init.testing.submodule"),
			Name:       "some-submodule",
			deployable: false,
		})
		t.Logf("Generating sub project for project: %v", testName)
		performProjectGeneration(subProjectData)
		execOrFail("./gradlew clean assemble check", testName, t)
	})
}

func TestMultiProjectGenerationDeployable(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	doTestOnEachLang(t, "multi_deployable", true, func(testName string, testLang languageTemplate) {
		if testLang.deployableConfig() == nil {
			t.Logf("Skipping non-deployable test: %v", testName)
			return
		}
		multiProjectData := makeDefaultProjectData(&ProjectData{
			Proj:  &multiProject{},
			Lang:  testLang,
			Group: packageName("com.g6init.testing"),
			Name:  "some-test-product",

			gdmcRepoURL: hackitGdmcRepo,
			deployable:  true,
		})

		t.Logf("Generating root project for project: %v", testName)
		performProjectGeneration(multiProjectData)
		execOrFail("./gradlew clean assemble check", testName, t)

		subProjectData := makeDefaultProjectData(&ProjectData{
			Proj:       &subProject{},
			Lang:       testLang,
			Group:      packageName("com.g6init.testing.submodule"),
			Name:       "some-submodule",
			deployable: true,
		})
		t.Logf("Generating sub project for project: %v", testName)
		performProjectGeneration(subProjectData)
		execOrFail("./gradlew clean assemble check", testName, t)
	})
}

func TestSingleProjectGenerationNoGdmcSimple(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	doTestOnEachLang(t, "single_nogdmc_simple", false, func(testName string, testLang languageTemplate) {
		data := makeDefaultProjectData(&ProjectData{
			Proj:  &singleProject{},
			Lang:  testLang,
			Group: packageName("com.g6init.testing"),
			Name:  "some-test-product",

			depResolver: mavenResolver,
			deployable:  false,
		})

		t.Logf("Generating project: %v", testName)
		performProjectGeneration(data)
		execOrFail("./gradlew clean assemble check", testName, t)
	})
}

func TestSingleProjectGenerationNoGdmcDeployable(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	doTestOnEachLang(t, "single_nogdmc_deployable", true, func(testName string, testLang languageTemplate) {
		if testLang.deployableConfig() == nil {
			t.Logf("Skipping non-deployable test: %v", testName)
			return
		}
		data := makeDefaultProjectData(&ProjectData{
			Proj:  &singleProject{},
			Lang:  testLang,
			Group: packageName("com.g6init.testing"),
			Name:  "some-test-product",

			depResolver: mavenResolver,
			deployable:  true,
		})

		t.Logf("Generating project: %v", testName)
		performProjectGeneration(data)
		execOrFail("./gradlew clean assemble check", testName, t)
	})
}

func TestMultiProjectGenerationNoGdmcSimple(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	doTestOnEachLang(t, "multi_nogdmc_simple", false, func(testName string, testLang languageTemplate) {
		multiProjectData := makeDefaultProjectData(&ProjectData{
			Proj:  &multiProject{},
			Lang:  testLang,
			Group: packageName("com.g6init.testing"),
			Name:  "some-test-product",

			depResolver: mavenResolver,
			deployable:  false,
		})

		t.Logf("Generating root project for project: %v", testName)
		performProjectGeneration(multiProjectData)
		execOrFail("./gradlew clean assemble check", testName, t)

		subProjectData := makeDefaultProjectData(&ProjectData{
			Proj:       &subProject{},
			Lang:       testLang,
			Group:      packageName("com.g6init.testing.submodule"),
			Name:       "some-submodule",
			deployable: false,
		})
		t.Logf("Generating sub project for project: %v", testName)
		performProjectGeneration(subProjectData)
		execOrFail("./gradlew clean assemble check", testName, t)
	})
}

func TestMultiProjectGenerationNoGdmcDeployable(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	doTestOnEachLang(t, "multi_nogdmc_deployable", true, func(testName string, testLang languageTemplate) {
		if testLang.deployableConfig() == nil {
			t.Logf("Skipping non-deployable test: %v", testName)
			return
		}
		multiProjectData := makeDefaultProjectData(&ProjectData{
			Proj:  &multiProject{},
			Lang:  testLang,
			Group: packageName("com.g6init.testing"),
			Name:  "some-test-product",

			depResolver: mavenResolver,
			deployable:  true,
		})

		t.Logf("Generating root project for project: %v", testName)
		performProjectGeneration(multiProjectData)
		execOrFail("./gradlew clean assemble check", testName, t)

		subProjectData := makeDefaultProjectData(&ProjectData{
			Proj:       &subProject{},
			Lang:       testLang,
			Group:      packageName("com.g6init.testing.submodule"),
			Name:       "some-submodule",
			deployable: true,
		})
		t.Logf("Generating sub project for project: %v", testName)
		performProjectGeneration(subProjectData)
		execOrFail("./gradlew clean assemble check", testName, t)
	})
}

func doTestOnEachLang(t *testing.T, testNamePrefix string, needsGit bool, testFunc func(testName string, testLang languageTemplate)) {
	startingDir := getwd()
	defer chdir(startingDir)

	workingDir := getAndPrepWorkingDir(startingDir)
	langs := reflect.ValueOf(projectLangs).MapKeys()
	for _, lng := range langs {
		testName := fmt.Sprintf("%v_%v", testNamePrefix, lng)
		t.Logf("Preparing project test: %v", testName)
		testLang := projectLangs[fmt.Sprint(lng)]
		dir := filepath.Join(workingDir, testName)
		prepAndChToProjectTestDir(dir, needsGit)
		testFunc(testName, testLang)
	}
}

func makeDefaultProjectData(data *ProjectData) *ProjectData {
	data.Version = defaultProjectVersion
	data.LicenseName = defaultLicenseName
	data.AndroidSdkDir = defaultAndroidSdkDir()
	data.AndroidNdkDir = defaultAndroidNdkDir()
	data.AndroidCompileSdkVersion = defaultAndroidCompileSdkVersion
	data.GradleVersion = defaultGradleVersion
	data.gitRepoURL = testingRepoURL
	return data
}

func getwd() string {
	startingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return startingDir
}

func getAndPrepWorkingDir(startingDir string) string {
	workingDir := filepath.Join(startingDir, "test_out")
	mkdir(workingDir)
	return workingDir
}

func prepAndChToProjectTestDir(dir string, needsGit bool) {
	err := os.RemoveAll(dir)
	if err != nil {
		panic(err)
	}
	mkdir(dir)
	chdir(dir)
	if needsGit {
		execOrPanic("git init")
		isGitRepoAsserted = false
	}
}

func chdir(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func execOrFail(command string, testName string, t *testing.T) string {
	t.Logf("Running command \"%v\" on project %v...", command, testName)
	output, err := exec.Command("bash", "-c", command).CombinedOutput()
	outputStr := string(output)
	if err != nil {
		t.Errorf("Failed to execute \"%v\" on project %v\nOuput:\n%v", command, testName, outputStr)
	} else {
		t.Logf("Succesfully ran \"%v\" on project %v\nOutput:\n%v", command, testName, outputStr)
	}
	return outputStr
}
