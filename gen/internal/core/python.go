package core

import (
	"encoding/json"
	"fmt"
	"gen/internal/core/commons"
	"gen/internal/core/py"
	"gen/internal/reader"
	"gen/internal/writer"
	"log"
)

type python struct {
	customNature
}

func (g *python) Init() Generator {
	g.files=make(map[string][]byte)
	g.model=make(map[string]interface{})
	g.spec=make([]byte,0)
	return g
}

func (g *python) WithName(name string) Generator {
	if name == ""{
		g.name=fmt.Sprintf("app%v", "test")
	}else{
		g.name=name
	}
	return g
}

func (g *python) WithPort(port int) Generator                          {
	g.port=port
	return g
}

func (g *python) WithInputSpec(spec interface{}) Generator {
	filePath := fmt.Sprintf("%v", spec)
	g.spec = reader.ReadSpec(filePath)
	err := json.Unmarshal([]byte(g.spec), &g.model)
	if err != nil {
		log.Println(err.Error())
	}
	return g
}

func (g *python) WithOutputPath(path string) Generator {
		g.outputPath = fmt.Sprintf("%v/%v/%v", path, "python", g.name)
	writer.CreateFolder(g.outputPath)
	specPath := fmt.Sprintf("%v/%v", g.outputPath, "spec")
	writer.CreateFolder(specPath)
	return g
}

func (g *python) Generate() {

	g.files["spec/spec.yml"] = commons.GenerateApiSpecFile(
		g.name,
		"api.",
		g.model)
	g.files["api.py"] = py.GenerateApi()
	g.files["server.py"] = py.GenerateServer(3001)
	g.files["requirements.txt"] = py.GenerateDependencies()

	for key, value := range g.files {
		filePath := fmt.Sprintf("%v/%v", g.outputPath, key)
		writer.WriteFile(filePath, value)
	}
}
