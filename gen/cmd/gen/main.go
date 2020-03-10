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
	//read and parse spec input as json file

	generator := core.NewGenerator(core.Custom)
	generator.WithTech("python")
	//write generatióon produt to output
	fmt.Println(outputDir)

}
