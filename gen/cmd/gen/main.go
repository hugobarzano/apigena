package main

import "fmt"

//func main() {
//
//	var inputSpec, outputDir string
//	flag.StringVar(&inputSpec, "input",
//		"", "input generation spec file")
//	flag.StringVar(&outputDir, "output",
//		"", "output directory to generated source-code")
//
//	flag.Parse()
//	fmt.Println("flags:", flag.Args())
//	fmt.Println(inputSpec)
//	//read and parse spec input as json file
//
//	generator := core.NewGenerator(core.Custom)
//	generator.WithTech("python")
//	//write generatióon produt to output
//	fmt.Println(outputDir)
//
//}

type ClientType int32

const (
	// DefaultClient defines bot-maincontroller client
	DefaultClient ClientType = iota
	// TechStackClient defines governance client
	TechStackClient
)

func getInterface() interface{} {
	return TechStackClient
}

func main()  {

	inter:=getInterface()
	fmt.Println(inter)
	switch inter {
	case TechStackClient:
		fmt.Printf("TECH")
	default:
		fmt.Printf("deafult")
	}


}