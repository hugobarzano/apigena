package core

type javascript struct{
	spec []byte
	files map[string][]byte
}

func (g *javascript)Init()Generator{
	g.files=make(map[string][]byte)
	return g
}

func (g *javascript) WithInputSpec(spec interface{}) Generator  {return g }


func (g *javascript) WithOutputPath(input string) Generator { return g}
func (g *javascript) Generate()  {}