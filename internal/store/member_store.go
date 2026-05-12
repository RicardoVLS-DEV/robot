package store

import (
	"context"
	"errors"
	"fmt"
	"robot/internal/domain"

	"github.com/jackc/pgx/v5"
)

type MemberStore struct {
	store *Store
}

func NewMemberStore(s *Store) *MemberStore {
	return &MemberStore{
		store: s,
	}
}

func (st *MemberStore) Insert(ctx context.Context, member *domain.Member, teamID domain.TeamID) (domain.MemberID, error) {
	query := `
		INSERT INTO member
		(name, email, is_leader, team_id)
		VALUES
		($1, $2, $3, $4)
		RETURNING id
	`

	var id domain.MemberID
	err := st.store.db.QueryRow(ctx, query, 
		member.Name,
		member.Email,
		member.IsLeader,
		teamID,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("intert member: %w", err)
	}

	return id, nil
}

func (st *MemberStore) FindByName(ctx context.Context, name string) (*domain.Member, domain.TeamID, error) {
	op := "FindByName"
	
	query := `
		SELECT id, name, email, is_leader, team_id
		FROM member
		WHERE name = $1
	`

	var id domain.MemberID
	var memberName, memberEmail string
	var isLeader bool
	var teamID domain.TeamID

	err := st.store.db.QueryRow(ctx, query, name).Scan(
		&id,
		&memberName,
		&memberEmail,
		&isLeader,
		&teamID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, 0, domain.NewRobotErr(op, "name", name, domain.ErrNotFound, "")
		}
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}

	member, err := domain.NewMember(memberName, memberEmail, isLeader)
	if err != nil {
		return nil, 0, err
	}
	member.ID = id
	return member, teamID, nil
}

func (st *MemberStore) FindByTeam(ctx context.Context, teamID domain.TeamID) ([]*domain.Member, error) {
	op := "FindByTeam"

	query := `
		SELECT id, name, email, is_leader
		FROM member
		WHERE team_id = $1
	`
	rows, err := st.store.db.Query(ctx, query, teamID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()
	
	var members []*domain.Member

	for rows.Next() {			
		var id domain.MemberID
		var memberName, memberEmail string
		var isLeader bool

		err := rows.Scan(&id, &memberName, &memberEmail, &isLeader)
		if err != nil {
			return nil, fmt.Errorf("%s: scan member: %w", op, err)
		}
		member, err := domain.NewMember(memberName, memberEmail, isLeader)
		if err != nil {
			return nil, err
		}
		member.ID = id
		members = append(members, member)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: read members: %w", op, err)
	}

	if len(members) == 0 {
		return nil, domain.NewRobotErr(op, "teamID", teamID, domain.ErrNotFound, "")
	}

	return members, nil
}

func (st *MemberStore) FindLeaderTeam(ctx context.Context, teamID domain.TeamID) (*domain.Member, error) {
	op := "FindLeaderTeam"

	query := `
		SELECT id, name, email, is_leader
		FROM member
		WHERE team_id = $1
		AND is_leader = true
	`
		
	var id domain.MemberID
	var memberName, memberEmail string
	var isLeader bool

	err := st.store.db.QueryRow(ctx, query, teamID).Scan(
		&id,
		&memberName,
		&memberEmail,
		&isLeader,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.NewRobotErr(op, "teamID", teamID, domain.ErrNotFound, "")
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	member, err := domain.NewMember(memberName, memberEmail, isLeader)
	if err != nil {
		return nil, err
	}
	member.ID = id
	return member, nil
}