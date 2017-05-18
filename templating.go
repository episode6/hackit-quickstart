package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"text/template"
)

type templatableConfig interface {
	templateAlias() string
}

func templateTemplateable(filename string, templateable templatableConfig, data interface{}) string {
	return templateAsset(templateable.templateAlias()+"/"+filename, data)
}

func templateTemplateableToFile(filename string, outputFilename string, templateable templatableConfig, data interface{}) {
	templateAssetToFile(templateable.templateAlias()+"/"+filename, outputFilename, data)
}

func templateAsset(filename string, data interface{}) string {
	var buffer bytes.Buffer
	templateAssetToWriter(filename, data, &buffer)
	return buffer.String()
}

func templateAssetToFile(filename string, outputFilename string, data interface{}) {
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	templateAssetToWriter(filename, data, outputFile)
}

func templateAssetToWriter(filename string, data interface{}, writer io.Writer) {
	tmplBytes, err := Asset("templates/" + filename)
	if err != nil {
		panic("ASSET NOT FOUND: " + filename)
	}

	templ := template.Must(template.New(filename).Parse(string(tmplBytes)))
	err = templ.Execute(writer, data)
	if err != nil {
		panic(err)
	}
}

func appendTextFile(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		panic(err)
	}
}

func writeRawTemplateAsset(assetPath string, destPath string, templatable templatableConfig) {
	writeRawAsset(templatable.templateAlias()+"/"+assetPath, destPath)
}

func writeRawAsset(assetPath string, destPath string) {
	bytes, err := Asset("templates/" + assetPath)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(destPath, bytes, 0755)
	if err != nil {
		panic(err)
	}
}
