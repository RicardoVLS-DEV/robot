package domain

type MemberID int

type Member struct {
	ID		MemberID
	Name 	string
	Email 	string
}

func NewMember(name, email string) (*Member, error) {
	if name == "" {
		return nil, ErrEmpty
	}

	return &Member{
		Name: name,
		Email: email,
	}, nil
}