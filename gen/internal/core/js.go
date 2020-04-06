package core

import (
	"encoding/json"
	"fmt"
	"gen/internal/core/commons"
	"gen/internal/core/js"
	"gen/internal/reader"
	"gen/internal/writer"
	"log"
)

type javascript struct{
	customNature
}

func (g *javascript)Init()Generator{
	g.files=make(map[string][]byte)
	g.model=make(map[string]interface{})
	g.spec=make([]byte,0)
	return g
}

func (g *javascript) WithName(name string) Generator                          {
	if name == ""{
		g.name=fmt.Sprintf("app%v", "test")
	}else{
		g.name=name
	}
	return g
}

func (g *javascript) WithPort(port int) Generator                          {
g.port=port
return g
}

func (g *javascript) WithInputSpec(spec interface{}) Generator  {
	filePath := fmt.Sprintf("%v", spec)
	g.spec = reader.ReadSpec(filePath)
	err := json.Unmarshal([]byte(g.spec), &g.model)
	if err != nil {
		log.Println(err.Error())
	}
	return g }


func (g *javascript) WithOutputPath(path string) Generator {
		g.outputPath = fmt.Sprintf("%v/%v/%v", path, "js", g.name)
	writer.CreateFolder(g.outputPath)
	specPath := fmt.Sprintf("%v/%v", g.outputPath, "spec")
	writer.CreateFolder(specPath)
	apiPath := fmt.Sprintf("%v/%v", g.outputPath, "api")
	writer.CreateFolder(apiPath)
	return g
}

func (g *javascript) Generate()  {

	g.files["spec/spec.yml"] = commons.GenerateApiSpecFile(
		g.name,
		"",
		g.model)
	g.files["api/index.js"] = js.GenerateApi()
	g.files["server.js"] = js.GenerateServer(g.port)
	g.files["package.json"] = js.GenerateDependencies(g.name)

	for key, value := range g.files {
		filePath := fmt.Sprintf("%v/%v", g.outputPath, key)
		writer.WriteFile(filePath, value)
	}
}