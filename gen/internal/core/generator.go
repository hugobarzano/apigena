package core

// Nature define generator nature
type Nature int32

const (
	// Base defines nothing
	Base Nature = iota
	// Go defines golang generator nature
	Go
	// Python
	Python
	// JS
	JS
)

// Generator interface definition
type Generator interface {
	Init()Generator
	WithName(name string) Generator
	WithPort(port int) Generator
	WithInputSpec(spec interface{}) Generator
	WithOutputPath(path string) Generator
	Generate()
}

type defaultNature struct{}
func (g *defaultNature) Init() Generator                          {return g}
func (g *defaultNature) WithName(name string) Generator           {return g}
func (g *defaultNature) WithPort(port int) Generator           {return g}
func (g *defaultNature) WithInputSpec(spec interface{}) Generator { return g }
func (g *defaultNature) WithOutputPath(path string) Generator     { return g }
func (g *defaultNature) Generate()                                {}
var DefaultNature = &defaultNature{}


type customNature struct {
	name string
	port int
	spec []byte
	files map[string][]byte
	outputPath string
	model      map[string]interface{}
}

// NewGenerator builds a generator from a given kind
func NewGenerator(nature Nature) Generator {
	switch nature {
	case Go:
		return &goApi{}
	case Python:
		return &python{}
	case JS:
		return &javascript{}
	default:
		return DefaultNature
	}
}

