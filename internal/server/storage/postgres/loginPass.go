package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/rs/zerolog"
)

const (
	lptCreate = "loginPassPostgres Create repo"
	lpUpdate  = "loginPassPostgres Update repo"
	lpDelete  = "loginPassPostgres Delete repo"
	lpGetAll  = "loginPassPostgres GetAll repo"
)

// LoginPassPostgres login-password data
type LoginPassPostgres struct {
	db  *sql.DB
	ver *Version
	log *zerolog.Logger
}

func NewLoginPassPostgres(db *sql.DB, ver *Version, log *zerolog.Logger) *LoginPassPostgres {
	return &LoginPassPostgres{
		db:  db,
		ver: ver,
		log: log,
	}
}

// Create creates a login-password data and updates data version
func (l *LoginPassPostgres) Create(ctx context.Context, lp *domain.LoginPass) (err error) {
	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		l.log.Err(err).Msgf("%s start tx error", lptCreate)
		return err
	}

	defer func() {
		if err != nil {
			txError := tx.Rollback()
			if txError != nil {
				l.log.Err(txError).Msgf("%s rollback error", lptCreate)
				err = fmt.Errorf("%s defer rollback error %s: %s", lptCreate, txError.Error(), err.Error())
			}
		}
	}()

	row := tx.QueryRowContext(ctx,
		"INSERT INTO public.login_pass(login, password, meta_info,user_id) VALUES ($1,$2,$3,$4) RETURNING id", lp.Login, lp.Password, lp.MetaInfo, lp.UserID)

	var id sql.NullInt64
	if err = row.Scan(&id); err != nil {
		l.log.Err(err).Msgf("%s scan id error", lptCreate)
		return err
	}

	if !id.Valid {
		err = fmt.Errorf("id lp no valid")
		l.log.Err(err).Msgf("%s error", lptCreate)
		return err
	}

	if err = l.ver.UpdateVersion(ctx, lp.UserID, tx); err != nil {
		l.log.Err(err).Msgf("%s error", lptCreate)
		return err
	}

	if err = tx.Commit(); err != nil {
		l.log.Err(err).Msgf("%s commit error", lptCreate)
		return err
	}

	lp.ID = int(id.Int64)

	return nil
}

// Update updates a login-password data and data version
func (l *LoginPassPostgres) Update(ctx context.Context, lp *domain.LoginPass) (err error) {
	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		l.log.Err(err).Msgf("%s start tx error", lpUpdate)
		return err
	}

	defer func() {
		if err != nil {
			txError := tx.Rollback()
			if txError != nil {
				l.log.Err(txError).Msgf("%s rollback error", lpUpdate)
				err = fmt.Errorf("%s defer rollback error %s: %s", lpUpdate, txError.Error(), err.Error())
			}
		}
	}()

	_, err = tx.ExecContext(ctx,
		"UPDATE public.login_pass SET login=$1, password=$2, meta_info=$3 WHERE id=$4 AND user_id=$5",
		lp.Login, lp.Password, lp.MetaInfo, lp.ID, lp.UserID)

	if err != nil {
		l.log.Err(err).Msgf("%s error", lpUpdate)
		return err
	}

	if err = l.ver.UpdateVersion(ctx, lp.UserID, tx); err != nil {
		l.log.Err(err).Msgf("%s error", lpUpdate)
		return err
	}

	if err = tx.Commit(); err != nil {
		l.log.Err(err).Msgf("%s commit error", lpUpdate)
		return err
	}

	return nil
}

// Delete deletes a login-password data and updates data version
func (l *LoginPassPostgres) Delete(ctx context.Context, lp *domain.LoginPass) (err error) {
	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		l.log.Err(err).Msgf("%s start tx error", lpDelete)
		return err
	}

	defer func() {
		if err != nil {
			txError := tx.Rollback()
			if txError != nil {
				l.log.Err(txError).Msgf("%s rollback error", lpDelete)
				err = fmt.Errorf("%s defer rollback error %s: %s", lpDelete, txError.Error(), err.Error())
			}
		}
	}()

	_, err = tx.ExecContext(ctx, "DELETE FROM public.login_pass WHERE id=$1 AND user_id=$2", lp.ID, lp.UserID)
	if err != nil {
		l.log.Err(err).Msgf("%s error", lpDelete)
		return err
	}

	if err = l.ver.UpdateVersion(ctx, lp.UserID, tx); err != nil {
		l.log.Err(err).Msgf("%s error", lpDelete)
		return err
	}

	if err = tx.Commit(); err != nil {
		l.log.Err(err).Msgf("%s commit error", lpDelete)
		return err
	}

	return nil
}

// GetAll gets all login-password data by user id
func (l *LoginPassPostgres) GetAll(ctx context.Context, userID int) ([]domain.LoginPass, error) {
	rows, err := l.db.QueryContext(ctx, "SELECT id, user_id, login, password, meta_info FROM login_pass WHERE user_id=$1", userID)
	if err != nil {
		l.log.Err(err).Msgf("%s get rows error", lpGetAll)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			l.log.Err(err).Msgf("%s rows close error", lpGetAll)
		}
	}()

	dataOutput := make([]domain.LoginPass, 0)
	for rows.Next() {
		var data domain.LoginPass
		rows.Scan(&data.ID, &data.UserID, &data.Login, &data.Password, &data.MetaInfo)

		dataOutput = append(dataOutput, data)
	}

	if err = rows.Err(); err != nil {
		l.log.Err(err).Msgf("%s rows scan error", lpGetAll)
	}

	return dataOutput, nil
}
