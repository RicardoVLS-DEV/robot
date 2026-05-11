package app

import (
	"robot/internal/domain"
	"robot/internal/excel"
)

func LoadFromExcel(path string) ([]*domain.Team, error) {
	parsedRows, err := excel.ParseRows(path)
	if err != nil {
		return nil, err
	}
	if err := buildFromParsedRows(parsedRows); err != nil {
		return nil, err
	}

	return nil, nil
}

func buildFromParsedRows(parsedRows []excel.FormRow) error {
	for _, row := range parsedRows {
		if err := validateRow("LoadFromExcel", row); err != nil {
			return err
		}
		if _, err := buildTeam(row); err != nil {
			return err
		}
	}
	return nil
}

func buildTeam(row excel.FormRow) (*domain.Team, error) {
	category, err := domain.NewCategory(row.Category, 0)
	if err != nil {
		return nil, err
	}

	Leader, err := domain.NewLeader(row.NameLeader, row.EmailLeader)
	if err != nil {
		return nil, err
	}

	members, err := buildMembers(row.Members, row.EmailMembers)
	if err != nil {
		return nil, err
	}

	var IDs []domain.MemberID

	for _, member := range members {
		IDs = append(IDs, member.ID)
	}

	return domain.NewTeam(
		row.NameTeam,
		row.School,
		row.Grade,
		IDs,
		Leader.ID,
		category.ID,
	)
}

func buildMembers(names, emails []string) ([]*domain.Member, error) {
	var members []*domain.Member

	for index, name := range names {
		email := ""
		if index < len(emails) {
			email = emails[index]
		}
		member, err := domain.NewMember(name, email)
		if err != nil {
			return members, err
		}
		members = append(members, member)
	}
	return members, nil
}

func validateRow(op string, rows excel.FormRow) error {
	if rows.NameTeam == "" {
		return createErr(op, "Name Team", domain.ErrEmpty)
	}

	if rows.NameLeader == "" {
		return createErr(op, "Name Leader", domain.ErrEmpty)
	}

	if rows.EmailLeader == "" {
		return createErr(op, "Email Leader", domain.ErrEmpty)
	}

	if rows.Category == "" {
		return createErr(op, "Category", domain.ErrEmpty)
	}

	if rows.School == "" {
		return createErr(op, "School", domain.ErrEmpty)
	}

	return nil
}

func createErr(op, field string, err error) *domain.RobotError {
	return domain.NewRobotErr(op, field, err)
}
