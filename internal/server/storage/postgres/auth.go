package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/kotche/gophKeeper/internal/server/domain"
	ers "github.com/kotche/gophKeeper/internal/server/domain/errs"
	"github.com/rs/zerolog"
)

const (
	authCreateUser = "authPostgres createUser repo"
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
	var userIdExist int
	row := a.db.QueryRowContext(ctx, "SELECT id FROM public.users WHERE login=$1", user.Username)
	err := row.Scan(&userIdExist)
	if err == nil {
		a.log.Warn().Msgf("%s conflict username '%s'", authCreateUser, user.Username)
		return ers.ConflictLoginError{
			Username: user.Username,
		}
	} else if err != nil && !errors.Is(err, sql.ErrNoRows) {
		a.log.Err(err).Msgf("%s scan userIdExist error", authCreateUser)
		return err
	}

	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		a.log.Err(err).Msgf("%s tx error", authCreateUser)
		return err
	}

	defer func() {
		if err != nil {
			txError := tx.Rollback()
			if txError != nil {
				a.log.Err(txError).Msgf("%s rollback error", authCreateUser)
				err = fmt.Errorf("%s defer rollback error %w: %s", authCreateUser, txError, err.Error())
			}
		}
	}()

	row = tx.QueryRowContext(ctx, "INSERT INTO public.users(login,password) VALUES ($1,$2) RETURNING id", user.Username, user.Password)
	var userID int
	if err = row.Scan(&userID); err != nil {
		a.log.Err(err).Msgf("%s scan userID error", authCreateUser)
		return err
	}

	if err = insertVersion(ctx, userID, tx, a.log); err != nil {
		a.log.Err(err).Msgf("%s error", authCreateUser)
		return err
	}

	if err = tx.Commit(); err != nil {
		a.log.Err(err).Msgf("%s commit error", authCreateUser)
		return err
	}

	user.ID = userID
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
