package core

import (
	"encoding/json"
	"fmt"
	"gen/internal/core/commons"
	"gen/internal/core/py"
	"gen/internal/reader"
	"io/ioutil"
	"log"
	"os"
)

//https://medium.com/@irshadhasmat/golang-simple-json-parsing-using-empty-interface-and-without-struct-in-go-language-e56d0e69968
type python struct {
	spec       []byte
	files      map[string][]byte
	model      map[string]interface{}
	outputPath string
}

func (g *python) Init() Generator {
	g.files = make(map[string][]byte)
	return g
}

func (g *python) WithInputSpec(spec interface{}) Generator {
	filePath := fmt.Sprintf("%v", spec)
	var err error
	g.spec, err = reader.ReadSpec(filePath)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal([]byte(g.spec), &g.model)
	if err != nil {
		log.Println(err.Error())
	}
	return g
}

func (g *python) WithOutputPath(input string) Generator {
		outPutPath:=fmt.Sprintf("%v/%v",input,"python")
		err:=os.MkdirAll(outPutPath, os.ModePerm)
		if err != nil{
			log.Println(err)
		}


	specPath:=fmt.Sprintf("%v/%v",outPutPath,"spec")
	err=os.MkdirAll(specPath, os.ModePerm)
	if err != nil{
		log.Println(err)
	}

	apiPath:=fmt.Sprintf("%v/%v",outPutPath,"api")
	err=os.MkdirAll(apiPath, os.ModePerm)
	if err != nil{
		log.Println(err)
	}

	g.outputPath = outPutPath

	return g
}
func (g *python) Generate() {

	g.files["spec/spec.yml"] = commons.GenerateApiSpecFile(
		"this is the T",
		"app",
		"App",
		"api.",
		g.model)

	g.files["server.py"]=py.GenerateServer(3001)
	g.files["api.py"]=py.GenerateApi()
	g.files["requirements.txt"]=py.GenerateDependencies()


	for key,value := range g.files{
		filePath:=fmt.Sprintf("%v/%v",g.outputPath,key)
		err := ioutil.WriteFile(filePath,value,os.FileMode(os.ModePerm))
		if err != nil {
			log.Fatal(err)
		}
	}
}


