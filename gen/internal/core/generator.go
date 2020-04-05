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
	WithInputSpec(spec interface{}) Generator
	WithOutputPath(path string) Generator
	Generate()
}

type base struct{}
func (g *base) Init() Generator                          {return g}
func (g *base) WithInputSpec(spec interface{}) Generator { return g }
func (g *base) WithOutputPath(path string) Generator     { return g }
func (g *base) Generate()                                {}



// DefaultNature
var DefaultNature = &base{}

// NewGenerator builds a generator from a given kind
func NewGenerator(nature Nature) Generator {
	switch nature {
	case Go:
		return &golang{}
	case Python:
		return &python{}
	case JS:
		return &javascript{}
	default:
		return DefaultNature
	}
}

