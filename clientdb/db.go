package clientdb

import (
	"context"
	"database/sql"
	"github.com/tsawler/goblender/client/clienthandlers/clientmodels"
	"time"
)

// DBModel holds the database
type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) GetPTMember(id int) (clientmodels.PTMember, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var s clientmodels.PTMember

	query := `select e.id, e.first_name, e.email, e.voted
			from pt_members e
			where e.id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&s.ID,
		&s.FirstName,
		&s.Email,
		&s.Voted,
	)
	if err != nil {
		return s, err
	}

	return s, nil
}

func (m *DBModel) GetFTMember(id int) (clientmodels.FTMember, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var s clientmodels.FTMember

	query := `select e.id, e.first_name, e.email, e.voted
			from ft_members e
			where e.id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&s.ID,
		&s.FirstName,
		&s.Email,
		&s.Voted,
	)
	if err != nil {
		return s, err
	}

	return s, nil
}
