package store

import (
	"context"
	"fmt"
	"robot/internal/domain"
)

type RobotStore struct {
	store *Store
}

func NewRobotStore(s *Store) *RobotStore {
	return &RobotStore{
		store: s,
	}
}

func (st *RobotStore) Insert(ctx context.Context, robot *domain.Robot, teamID domain.TeamID) (domain.RobotID, error) {
	query := `
		INSERT INTO robot
		(name, weight_cm, width_cm, height_cm, length_cm,  is_valid, invalid_reason, autonomous, power_button, internal_power, status,  team_id)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id
	`

	var id domain.RobotID
	err := st.store.db.QueryRow(ctx, query,
		robot.Name,
		robot.Weight,
		robot.Width,
		robot.Height,
		robot.Length,
		robot.IsValid,
		robot.InvalidReason,
		robot.Autonomous,
		robot.PowerButton,
		robot.InternalPower,
		robot.Status,

		teamID,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("intert robott: %w", err)
	}

	return id, nil
}

func (st *RobotStore) FindByTeamRobot(ctx context.Context, teamID domain.TeamID) ([]*domain.Robot, error) {
	op := "FindByTeamRobot"

	query := `
		SELECT id, name, weight_cm, width_cm, height_cm, length_cm, is_valid, invalid_reason, autonomous, power_button, internal_power, status, team_id
		FROM robot
		WHERE team_id = $1
	`

	rows, err := st.store.db.Query(ctx, query, teamID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var robots []*domain.Robot

	for rows.Next() {
		var id domain.RobotID
		var robotName, robotInvalidReason, robotAutonomous, robotPowerButton, robotInternalPower, robotStatus string
		var robotWeightCm, robotWidthCm, robotHeightCm, robotLengthCm float64
		var isValid bool
		var robotTeamID domain.TeamID

		err := rows.Scan(
			&id,
			&robotName,
			&robotWeightCm,
			&robotWidthCm,
			&robotHeightCm,
			&robotLengthCm,
			&isValid,
			&robotInvalidReason,
			&robotAutonomous,
			&robotPowerButton,
			&robotInternalPower,
			&robotStatus,
			&robotTeamID,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: scan robot: %w", op, err)
		}

		robot, err := domain.NewRobot(
			robotName,
			robotWeightCm,
			robotWidthCm,
			robotHeightCm,
			robotLengthCm,
			robotAutonomous,
			robotPowerButton,
			robotInternalPower,
			robotTeamID,
		)
		if err != nil {
			return nil, err
		}

		robot.ID = id
		robot.IsValid = isValid
		robot.InvalidReason = robotInvalidReason
		robot.Status = robotStatus

		robots = append(robots, robot)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: read robots: %w", op, err)
	}

	if len(robots) == 0 {
		return nil, domain.NewRobotErr(op, "teamID", teamID, domain.ErrNotFound, "")
	}

	return robots, nil
}
