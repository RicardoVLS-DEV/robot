package domain

type Integrante struct {
	Name string
	Email string
	Lider bool
}

func NewIntegrante(name, email string, lider bool) (*Integrante, error) {
	if name == "" {
		return nil, ErrEmpty
	}

	return &Integrante{
		Name: name,
		Email: email,
		Lider: lider,
	}, nil
}