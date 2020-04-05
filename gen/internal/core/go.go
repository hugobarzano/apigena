package core

import (
	"fmt"
	"gen/internal/reader"
	"github.com/ChimeraCoder/gojson"
	"log"
)

type golang struct{
customNature
	model      []byte
}

func (g *golang)Init()Generator{
	g.files=make(map[string][]byte)
	return g
}

func (g *golang) WithName(name string) Generator                          {
	g.name=name
	return g
	}

func (g *golang) WithPort(port int) Generator                          {
	g.port=port
	return g
}


func (g *golang) WithInputSpec(spec interface{}) Generator  {
	filePath:= fmt.Sprintf("%v",spec)
	specReader,err:=reader.GetReader(filePath)
	if err!=nil{
		log.Println(err.Error())
	}
	g.model, err = gojson.Generate(
		specReader,
		gojson.ParseJson,
		"test",
		"main",[]string{"json","yml"},
		true, true)
	if err != nil{
		log.Println(err.Error())
	}
	return g
	}

func (g *golang) WithOutputPath(input string) Generator { return g}
func (g *golang) Generate()  {}
