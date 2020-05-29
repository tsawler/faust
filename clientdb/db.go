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

func (m *DBModel) VoteYesFT(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update vote_totals set yes = yes + 1 where id = 1`
	_, err := m.DB.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	query = `update ft_members set voted = 1 where id = $1`
	_, err = m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) VoteNoFT(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update vote_totals set no = no + 1 where id = 1`
	_, err := m.DB.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	query = `update ft_members set voted = 1 where id = $1`
	_, err = m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) VoteYesPT(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update vote_totals set yes = yes + 1 where id = 2`
	_, err := m.DB.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	query = `update pt_members set voted = 1 where id = $1`
	_, err = m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) VoteNoPT(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update vote_totals set no = no + 1 where id = 2`
	_, err := m.DB.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	query = `update pt_members set voted = 1 where id = $1`
	_, err = m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
