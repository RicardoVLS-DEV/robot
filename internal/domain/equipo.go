package domain

type Equipo struct {
	Name 			string
	Integrantes 	[]*Integrante
	Escuela 		string
	Nivel 			string
	Resultado 		string
}

func NewEquipo(name, escuela, nivel string, integrantes []*Integrante) (*Equipo, error) {
	return &Equipo{
		Name: name,
		Integrantes: integrantes,
		Escuela: escuela,
		Nivel: nivel,
	}, nil
}