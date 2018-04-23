package mavenutil

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	version "github.com/hashicorp/go-version"
)

// MavenResolver is an object that can resolve versions for maven dependencies
type MavenResolver map[string]string

type metadata struct {
	Version    string             `xml:"version"`
	Versioning metadataVersioning `xml:"versioning"`
}

type metadataVersioning struct {
	Versions metadataVersions `xml:"versions"`
}

type metadataVersions struct {
	Versions []string `xml:"version"`
}

var mavenRepos = []string{
	"https://jcenter.bintray.com/",
	"https://maven.google.com/",
}

// FormatKeys formats a slice of maven dependency keys and applies versions to them
func (mr *MavenResolver) FormatKeys(keys []string) []string {
	formattedDeps := make([]string, len(keys))
	for i, k := range keys {
		formattedDeps[i] = mr.formatKey(k)
	}
	return formattedDeps
}

func (mr *MavenResolver) formatKey(key string) string {
	return fmt.Sprintf("%v:%v", key, mr.FindVersion(key))
}

// FindVersion returns just the version for the provided key
func (mr *MavenResolver) FindVersion(key string) string {
	cachedVersionMap := *mr
	cachedVersion := cachedVersionMap[key]
	if cachedVersion != "" {
		return cachedVersion
	}

	foundVersion := ""
	for _, mavenURL := range mavenRepos {
		foundVersion = highestVersion(foundVersion, mr.findBestVersionFromRepo(key, mavenURL))
	}
	if foundVersion == "" {
		panic("Could not find version for key: " + key)
	}
	cachedVersionMap[key] = foundVersion
	return foundVersion
}

func (mr *MavenResolver) findBestVersionFromRepo(key string, metadataURLBase string) string {
	keyAsPath := strings.Replace(key, ":", "/", -1)
	keyAsPath = strings.Replace(keyAsPath, ".", "/", -1)
	metadataURL := metadataURLBase + keyAsPath + "/maven-metadata.xml"
	resp, err := http.Get(metadataURL)
	if err != nil {
		panic("Could not load url: " + metadataURL)
	}
	defer resp.Body.Close()

	xmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Could not read data for url: " + metadataURL)
	}

	var md metadata
	xml.Unmarshal(xmlData, &md)
	return md.findBestVersion()
}

func (md *metadata) findBestVersion() string {
	// we prefer versions without weird text in them, but will fallback if we must
	regx := regexp.MustCompile(`^[0-9.]*$`)
	for i := len(md.Versioning.Versions.Versions) - 1; i >= 0; i-- {
		v := md.Versioning.Versions.Versions[i]
		if regx.MatchString(v) {
			return v
		}
	}
	return md.Version
}

func highestVersion(version1 string, version2 string) string {
	if version1 == "" {
		return version2
	}
	if version2 == "" {
		return version1
	}
	v1, err := version.NewVersion(version1)
	if err != nil {
		return ""
	}
	v2, err := version.NewVersion(version2)
	if err != nil {
		return ""
	}
	if v1.GreaterThan(v2) {
		return version1
	}
	return version2
}
