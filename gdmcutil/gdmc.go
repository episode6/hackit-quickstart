package gdmcutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// GdmcDep is a representation of a single gdmc dependency
type GdmcDep struct {
	GroupID        string `json:"groupId"`
	ArtifactID     string `json:"artifactId"`
	Version        string `json:"version"`
	InheritVersion string `json:"inheritVersion"`
}

// GdmcMap is a representation of an entrie gdmc json file
type GdmcMap map[string]*json.RawMessage

// LoadMap loads a new GdmcMap from the provided json file
func LoadMap(gdmcFilename string) *GdmcMap {
	gdmcFile, err := ioutil.ReadFile(gdmcFilename)
	if err != nil {
		panic(err)
	}

	var gdmcMap GdmcMap
	json.Unmarshal(gdmcFile, &gdmcMap)
	return &gdmcMap
}

// GetDep returns the GdmcDep for the given key
func (gdmcMap *GdmcMap) GetDep(depKey string) GdmcDep {
	var dep GdmcDep
	gm := *gdmcMap
	rawJSON := gm[depKey]
	json.Unmarshal(*rawJSON, &dep)
	return dep
}

// FormatKey takes the supplied gdmc key and returns
// the same dependency in a gradle-acceptable format with
// the version included (i.e. group:name:version)
func (gdmcMap *GdmcMap) FormatKey(depKey string) string {
	version := gdmcMap.FindVersion(depKey)
	return fmt.Sprintf("%v:%v", depKey, version)
}

//FindVersion returns just the version for the supplied key
func (gdmcMap *GdmcMap) FindVersion(key string) string {
	dep := gdmcMap.GetDep(key)
	if dep.Version != "" {
		return dep.Version
	}
	if dep.InheritVersion != "" {
		return gdmcMap.FindVersion(dep.InheritVersion)
	}
	panic("Could not find version for key: " + key)
}

// FormatKeys executes FormatKey on a slice of keys
func (gdmcMap *GdmcMap) FormatKeys(keys []string) []string {
	formattedDeps := make([]string, len(keys))
	for i, k := range keys {
		formattedDeps[i] = gdmcMap.FormatKey(k)
	}
	return formattedDeps
}
