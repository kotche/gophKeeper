package postgres

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"
)

// insertVersion adds a version of the data by user ID
func insertVersion(ctx context.Context, userID int, tx *sql.Tx, l *zerolog.Logger) error {
	_, err := tx.ExecContext(ctx, "INSERT INTO public.versions(version, user_id) VALUES ($1,$2)", 0, userID)
	if err != nil {
		l.Err(err).Msg("insertVersion error")
		return err
	}
	return nil
}

// updateVersion updates the data version by user ID
func updateVersion(ctx context.Context, userID int, tx *sql.Tx, l *zerolog.Logger) error {
	_, err := tx.ExecContext(ctx, "UPDATE public.versions SET version = version + 1 WHERE user_id = $1", userID)
	if err != nil {
		l.Err(err).Msg("updateVersion error")
		return err
	}
	return nil
}
