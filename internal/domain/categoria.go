package domain

type Categoria struct {
	Name string
	Equipo 	[]*Equipo
}

func NewCategoria(name string, equipo *Equipo) (*Categoria, error) {
	return &Categoria{
		Name: name,
		Equipo: []*Equipo{equipo},
	}, nil
}

func (c *Categoria) AddEquipo(equipo *Equipo) error {
	c.Equipo = append(c.Equipo, equipo)
	return nil
}
