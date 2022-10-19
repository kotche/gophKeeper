package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog"
)

// Version data version db
type Version struct {
	db  *sql.DB
	log *zerolog.Logger
}

func NewVersionData(db *sql.DB, log *zerolog.Logger) *Version {
	return &Version{
		db:  db,
		log: log,
	}
}

// InsertVersion adds a version of the data by user ID
func (v *Version) InsertVersion(ctx context.Context, userID int, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "INSERT INTO public.versions(version, user_id) VALUES ($1,$2)", 0, userID)
	if err != nil {
		v.log.Err(err).Msg("insertVersion error")
		return err
	}
	return nil
}

// UpdateVersion updates the data version by user ID
func (v *Version) UpdateVersion(ctx context.Context, userID int, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "UPDATE public.versions SET version = version + 1 WHERE user_id = $1", userID)
	if err != nil {
		v.log.Err(err).Msg("updateVersion error")
		return err
	}
	return nil
}

// GetVersion get the current version of the data
func (v *Version) GetVersion(ctx context.Context, userID int) (uint, error) {
	row := v.db.QueryRowContext(ctx, "SELECT version FROM public.versions WHERE user_id=$1", userID)
	var version uint
	if err := row.Scan(&version); err != nil {
		v.log.Err(err).Msg("commonsPostgres GetVersion repo scan version error")
		return 0, errors.New("get version error")
	}
	return version, nil
}
