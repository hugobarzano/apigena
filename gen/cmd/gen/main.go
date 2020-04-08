package main

import (
	"flag"
	"fmt"
	"gen/internal/core"
	"log"
)

func main() {

	var name, inputSpec, outputDir string

	flag.StringVar(&name, "name",
		"", "Application name to be generated")
	port := flag.Int("port", 3000, "TCP port of generated application")
	tech := flag.String("tech", "go", "Technology to generate: go, python, js")
	flag.StringVar(&inputSpec, "input",
		"", "Input generation spec file (json) format")
	flag.StringVar(&outputDir, "output",
		"", "Output directory to generated source-code")

	flag.Parse()
	if inputSpec == "" {
		log.Fatalf("-input is required")
	}
	if outputDir == "" {
		log.Fatalf("-output is required")
	}

	var nature core.Nature
	switch *tech {
	case "python":
		nature = core.Python
	case "js":
		nature = core.JS
	case "go":
		nature = core.Go
	default:
		log.Fatalf(fmt.Sprintf("Technology %v not supported.", *tech))
	}

	generator := core.New(nature)
	generator.Init().
		WithName(name).
		WithPort(*port).
		WithInputSpec(inputSpec).
		WithOutputPath(outputDir).Generate()

	log.Println(fmt.Sprintf("Checkout %v ", outputDir))
}
