package core

import (
	"encoding/json"
	"fmt"
	"gen/internal/core/commons"
	"gen/internal/core/golang"
	"gen/internal/reader"
	"gen/internal/writer"
	"github.com/ChimeraCoder/gojson"
	"log"
	"strings"
)

type goApi struct {
	customNature
	nativeModel []byte
}

func (g *goApi) Init() Generator {
	g.files = make(map[string][]byte)
	g.model = make(map[string]interface{})
	g.nativeModel = make([]byte, 0)
	g.spec = make([]byte, 0)
	return g
}

func (g *goApi) WithName(name string) Generator {
	if name == "" {
		g.name = fmt.Sprintf("app%v", "test")
	} else {
		g.name = name
	}
	return g
}

func (g *goApi) WithPort(port int) Generator {
	g.port = port
	return g
}

func (g *goApi) WithInputSpec(spec interface{}) Generator {
	filePath := fmt.Sprintf("%v", spec)
	g.spec = reader.ReadSpec(filePath)
	err := json.Unmarshal([]byte(g.spec), &g.model)
	if err != nil {
		log.Println(err.Error())
	}

	specReader, err := reader.GetReader(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	modelGenerated, err := gojson.Generate(
		specReader,
		gojson.ParseJson,
		strings.ToUpper(g.name),
		"model", []string{"json", "yml"},
		true, true)
	if err != nil {
		log.Println(err.Error())
	}

	g.nativeModel = customizeModel(modelGenerated)
	return g
}

func customizeModel(inputModel []byte) []byte {
	stringModel := string(inputModel)
	lastIndex := strings.LastIndex(stringModel, "}")
	IDField := "    ID string      `json:\"id\" yml:\"id\"` \n}"
	return []byte(stringModel[:lastIndex] + IDField + stringModel[lastIndex+1:])
}

func (g *goApi) WithOutputPath(path string) Generator {

	g.outputPath = fmt.Sprintf("%v/%v/%v", path, "go", g.name)
	writer.CreateFolder(g.outputPath)
	specPath := fmt.Sprintf("%v/%v", g.outputPath, "spec")
	writer.CreateFolder(specPath)
	apiPath := fmt.Sprintf("%v/%v", g.outputPath, "api")
	writer.CreateFolder(apiPath)
	modelPath := fmt.Sprintf("%v/%v", g.outputPath, "model")
	writer.CreateFolder(modelPath)
	templatesPath := fmt.Sprintf("%v/%v", g.outputPath, "templates")
	writer.CreateFolder(templatesPath)
	return g
}
func (g *goApi) Generate() {

	data := make(map[string]interface{})
	data["api"] = strings.ToLower(g.name)
	data["model"] = strings.ToUpper(g.name)
	data["port"] = g.port

	g.files["spec/spec.yml"] = commons.GenerateApiSpecFile(
		g.name,
		"api.",
		g.model)
	g.files["templates/index.html"] = commons.GenerateIndex("")
	g.files["api/api.go"] = golang.GenerateApi(data)
	g.files["model/model.go"] = g.nativeModel
	g.files["server.go"] = golang.GenerateServer(data)
	g.files["go.mod"] = golang.GenerateDependencies(data)

	for key, value := range g.files {
		filePath := fmt.Sprintf("%v/%v", g.outputPath, key)
		writer.WriteFile(filePath, value)
	}
}
