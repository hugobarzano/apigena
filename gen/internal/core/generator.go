package core

// Nature define generator nature
type Nature int32

const (
	// Base defines nothing
	Base Nature = iota
	// Custom defines psi notifier
	Custom
)

// BaseGenerator
var BaseGenerator = &base{}

// Generator interface definition
type Generator interface {
	WithTech(tech string) Generator
}

type base struct{}
func (g *base) WithTech(tech string) Generator  { return g }


type custom struct{
	tech string
	spec map[string]interface{}
}
func (g *custom) WithTech(tech string) Generator  {
	g.tech=tech
	return g }

// NewGenerator builds a generator from a given kind
func NewGenerator(nature Nature) Generator {
	switch nature {
	case Custom:
		return &custom{
			spec: map[string]interface{}{},
		}
	default:
		return BaseGenerator
	}
}

