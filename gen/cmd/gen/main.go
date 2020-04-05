package main

import (
	"flag"
	"fmt"
	"gen/internal/core"
)

func main() {

	var inputSpec, outputDir string
	flag.StringVar(&inputSpec, "input",
		"", "input generation spec file")
	flag.StringVar(&outputDir, "output",
		"", "output directory to generated source-code")

	flag.Parse()
	fmt.Println("flags:", flag.Args())
	fmt.Println(inputSpec)
	fmt.Println(outputDir)


	pythonGene :=core.NewGenerator(core.Python)
	pythonGene.Init().
		WithInputSpec(inputSpec).
		WithOutputPath(outputDir).Generate()
}

