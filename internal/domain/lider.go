package domain

type LeaderID int

type Leader struct {
	ID    LeaderID
	Name  string
	Email string
}

func NewLeader(name, email string) (*Leader, error) {
	if name == "" || email == "" {
		return nil, ErrEmpty
	}
	return &Leader{
		Name:  name,
		Email: email,
	}, nil
}
