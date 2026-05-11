package store

import "robot/internal/domain"

func (s *Store) InsertLeader(Leader *domain.Leader) (domain.LeaderID, error) {
	var id domain.LeaderID

	row := s.db.QueryRow(
		`INSERT INTO Leaders 
		(name, email)
		VALUES
		($1, $2)`,
		Leader.Name, Leader.Email,
	)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
