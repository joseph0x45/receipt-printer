package repository

import (
	"app/models"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SessionRepo struct {
	db *sqlx.DB
}

func NewSessionRepo(
	db *sqlx.DB,
) *SessionRepo {
	return &SessionRepo{
		db: db,
	}
}

func (r *SessionRepo) Insert(session *models.Session) error {
	const query = `
    insert into sessions(
      id, active
    )
    values (
      :id, :active
    )
  `
	_, err := r.db.NamedExec(query, session)
	if err != nil {
		return fmt.Errorf("Error while inserting session: %w", err)
	}
	return nil
}

func (r *SessionRepo) GetByID(id string) (*models.Session, error) {
	const query = "select * from sessions where id=$1 and active=true"
	session := &models.Session{}
	err := r.db.Get(session, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("Error while getting session: %w", err)
	}
	return session, nil
}
