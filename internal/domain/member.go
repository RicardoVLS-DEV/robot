package domain

type MemberID int

type Member struct {
	ID		MemberID
	Name 	string
	Email 	string
	IsLeader bool
}

func NewMember(name, email string, isLeader bool) (*Member, error) {
	if name == "" {
		return nil, ErrEmpty
	}

	return &Member{
		Name: name,
		Email: email,
		IsLeader: isLeader,
	}, nil
}