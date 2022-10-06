package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kotche/gophKeeper/internal/server/domain"
	ers "github.com/kotche/gophKeeper/internal/server/domain/errs"
	"github.com/rs/zerolog"
)

type AuthPostgres struct {
	db  *sql.DB
	log *zerolog.Logger
}

func NewAuthPostgres(db *sql.DB, log *zerolog.Logger) *AuthPostgres {
	return &AuthPostgres{
		db:  db,
		log: log,
	}
}

func (a *AuthPostgres) CreateUser(ctx context.Context, user *domain.User) error {
	const fInfo = "authPostgres createUser repo"
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		a.log.Err(err).Msgf("%s tx error", fInfo)
		return err
	}

	defer func() {
		if err != nil {
			txError := tx.Rollback()
			if txError != nil {
				a.log.Err(txError).Msgf("%s rollback error", fInfo)
				err = fmt.Errorf("%s defer rollback error %w: %s", fInfo, txError, err.Error())
			}
		}
	}()

	row := tx.QueryRowContext(ctx, "INSERT INTO public.users(login,password) VALUES ($1,$2) RETURNING id", user.Username, user.Password)
	var userID sql.NullInt64
	if err = row.Scan(&userID); err != nil {
		a.log.Err(err).Msgf("%s scan userID error", fInfo)
	}
	if !userID.Valid {
		return ers.ConflictLoginError{
			Username: user.Username,
		}
	}

	if err = insertVersion(ctx, int(userID.Int64), tx, a.log); err != nil {
		a.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = tx.Commit(); err != nil {
		a.log.Err(err).Msgf("%s commit error", fInfo)
		return err
	}

	user.ID = int(userID.Int64)
	return nil
}

func (a *AuthPostgres) GetUserID(ctx context.Context, user *domain.User) (int, error) {
	row := a.db.QueryRowContext(ctx, "SELECT id FROM public.users WHERE login=$1 AND password=$2", user.Username, user.Password)
	var output sql.NullInt64
	if err := row.Scan(&output); err != nil {
		a.log.Err(err).Msg("authPostgres getUserID repo scan users user_id error")
	}
	if !output.Valid {
		return 0, ers.AuthenticationError{}
	}
	userID := int(output.Int64)
	return userID, nil
}
