package domain

type TeamID int

type Team struct {
	ID         TeamID
	Name       string
	School     string
	Grade      string
	Result     string
	MembersIDs []MemberID
	LeaderID   LeaderID
	CategoryID CategoryID
}

func NewTeam(name, school, grade string, membersID []MemberID, LeaderID LeaderID, categoryID CategoryID) (*Team, error) {
	if name == "" || school == "" || grade == "" {
		return nil, ErrEmpty
	}

	if len(membersID) < 1 {
		return nil, ErrNotEnoughMembers
	}

	return &Team{
		Name:       name,
		School:     school,
		Grade:      grade,
		MembersIDs: membersID,
		LeaderID:   LeaderID,
		CategoryID: categoryID,
	}, nil
}
