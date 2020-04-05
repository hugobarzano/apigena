package main

import (
	"flag"
	"fmt"
	"gen/internal/core"
)

func main() {

	var name,inputSpec, outputDir string
	flag.StringVar(&name, "name",
		"", "application name to be generated")
	flag.StringVar(&inputSpec, "input",
		"", "input generation spec file")
	flag.StringVar(&outputDir, "output",
		"", "output directory to generated source-code")

	flag.Parse()
	fmt.Println("flags:", flag.Args())
	fmt.Println(inputSpec)
	fmt.Println(outputDir)
	fmt.Println(name)


	pythonGene :=core.NewGenerator(core.Python)
	pythonGene.Init().
		WithName(name).
		WithPort(3001).
		WithInputSpec(inputSpec).
		WithOutputPath(outputDir).Generate()


	JsGene :=core.NewGenerator(core.JS)
	JsGene.Init().
		WithName(name).
		WithPort(3002).
		WithInputSpec(inputSpec).
		WithOutputPath(outputDir).Generate()
}

