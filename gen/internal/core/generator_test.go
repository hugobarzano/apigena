package core

import (
	"fmt"
	"github.com/ChimeraCoder/gojson"
	"strings"
	"testing"
)

func TestGen(t *testing.T) {


	file := strings.NewReader(`{"foo" : "bar", "baz" : null}`)
	model, err := gojson.Generate(file, gojson.ParseJson, "Cat", "model", []string{"json","yml"}, true, true)

	if err !=nil{
		fmt.Println(err.Error())
	}

	fmt.Println(string(model))
}