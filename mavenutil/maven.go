package mavenutil

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
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

// FormatKeys formats a slice of maven dependency keys and applies versions to them
func (mr *MavenResolver) FormatKeys(keys []string) []string {
	formattedDeps := make([]string, len(keys))
	for i, k := range keys {
		formattedDeps[i] = mr.formatKey(k)
	}
	return formattedDeps
}

func (mr *MavenResolver) formatKey(key string) string {
	cachedVersionMap := *mr
	cachedVersion := cachedVersionMap[key]
	if cachedVersion != "" {
		return fmt.Sprintf("%v:%v", key, cachedVersion)
	}

	keyAsPath := strings.Replace(key, ":", "/", -1)
	keyAsPath = strings.Replace(keyAsPath, ".", "/", -1)
	metadataURL := "https://jcenter.bintray.com/" + keyAsPath + "/maven-metadata.xml"
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
	version := md.findBestVersion()
	if version == "" {
		panic("Could not find version for key: " + key)
	}
	cachedVersionMap[key] = version
	return fmt.Sprintf("%v:%v", key, version)
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
