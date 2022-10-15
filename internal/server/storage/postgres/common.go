package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog"
)

// CommonsPostgres common repository
type CommonsPostgres struct {
	db  *sql.DB
	log *zerolog.Logger
}

func NewCommonPostgres(db *sql.DB, log *zerolog.Logger) *CommonsPostgres {
	return &CommonsPostgres{
		db:  db,
		log: log,
	}
}

// GetVersion get the current version of the data
func (c *CommonsPostgres) GetVersion(ctx context.Context, userID int) (uint, error) {
	row := c.db.QueryRowContext(ctx, "SELECT version FROM public.versions WHERE user_id=$1", userID)
	var version uint
	if err := row.Scan(&version); err != nil {
		c.log.Err(err).Msg("commonsPostgres GetVersion repo scan version error")
		return 0, errors.New("get version error")
	}
	return version, nil
}
