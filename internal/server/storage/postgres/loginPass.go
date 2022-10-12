package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/rs/zerolog"
)

// LoginPassPostgres login-password pair data
type LoginPassPostgres struct {
	db  *sql.DB
	log *zerolog.Logger
}

func NewLoginPassPostgres(db *sql.DB, log *zerolog.Logger) *LoginPassPostgres {
	return &LoginPassPostgres{
		db:  db,
		log: log,
	}
}

// Create creates a login-password pair and updates data version
func (l *LoginPassPostgres) Create(ctx context.Context, lp *domain.LoginPass) (err error) {
	const fInfo = "loginPassPostgres create repo"

	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		l.log.Err(err).Msgf("%s start tx error", fInfo)
		return err
	}

	defer func() {
		if err != nil {
			txError := tx.Rollback()
			if txError != nil {
				l.log.Err(txError).Msgf("%s rollback error", fInfo)
				err = fmt.Errorf("%s defer rollback error %s: %s", fInfo, txError.Error(), err.Error())
			}
		}
	}()

	row := tx.QueryRowContext(ctx,
		"INSERT INTO public.login_pass(login, password, meta_info,user_id) VALUES ($1,$2,$3,$4) RETURNING id", lp.Login, lp.Password, lp.MetaInfo, lp.UserID)

	var id sql.NullInt64
	if err = row.Scan(&id); err != nil {
		l.log.Err(err).Msgf("%s scan id error", fInfo)
		return err
	}

	if !id.Valid {
		err = fmt.Errorf("id lp no valid")
		l.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = updateVersion(ctx, lp.UserID, tx, l.log); err != nil {
		l.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = tx.Commit(); err != nil {
		l.log.Err(err).Msgf("%s commit error", fInfo)
		return err
	}

	lp.ID = int(id.Int64)

	return nil
}

// Update updates a login-password pair and data version
func (l *LoginPassPostgres) Update(ctx context.Context, lp *domain.LoginPass) (err error) {
	const fInfo = "loginPassPostgres update repo"

	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		l.log.Err(err).Msgf("%s start tx error", fInfo)
		return err
	}

	defer func() {
		if err != nil {
			txError := tx.Rollback()
			if txError != nil {
				l.log.Err(txError).Msgf("%s rollback error", fInfo)
				err = fmt.Errorf("%s defer rollback error %s: %s", fInfo, txError.Error(), err.Error())
			}
		}
	}()

	_, err = tx.ExecContext(ctx,
		"UPDATE public.login_pass SET login=$1, password=$2, meta_info=$3 WHERE id=$4 AND user_id=$5",
		lp.Login, lp.Password, lp.MetaInfo, lp.ID, lp.UserID)

	if err != nil {
		l.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = updateVersion(ctx, lp.UserID, tx, l.log); err != nil {
		l.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = tx.Commit(); err != nil {
		l.log.Err(err).Msgf("%s commit error", fInfo)
		return err
	}

	return nil
}

// Delete deletes a login-password pair and updates data version
func (l *LoginPassPostgres) Delete(ctx context.Context, lp *domain.LoginPass) (err error) {
	const fInfo = "loginPassPostgres delete repo"

	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		l.log.Err(err).Msgf("%s start tx error", fInfo)
		return err
	}

	defer func() {
		if err != nil {
			txError := tx.Rollback()
			if txError != nil {
				l.log.Err(txError).Msgf("%s rollback error", fInfo)
				err = fmt.Errorf("%s defer rollback error %s: %s", fInfo, txError.Error(), err.Error())
			}
		}
	}()

	_, err = tx.ExecContext(ctx, "DELETE FROM public.login_pass WHERE id=$1 AND user_id=$2", lp.ID, lp.UserID)
	if err != nil {
		l.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = updateVersion(ctx, lp.UserID, tx, l.log); err != nil {
		l.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = tx.Commit(); err != nil {
		l.log.Err(err).Msgf("%s commit error", fInfo)
		return err
	}

	return nil
}

// GetAll returns all login-password pairs by user id
func (l *LoginPassPostgres) GetAll(ctx context.Context, userID int) ([]domain.LoginPass, error) {
	const fInfo = "loginPassPostgres getAll repo"

	rows, err := l.db.QueryContext(ctx, "SELECT id, user_id, login, password, meta_info FROM login_pass WHERE user_id=$1", userID)
	if err != nil {
		l.log.Err(err).Msgf("%s get rows error", fInfo)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			l.log.Err(err).Msgf("%s rows close error", fInfo)
		}
	}()

	dataOutput := make([]domain.LoginPass, 0)
	for rows.Next() {
		var data domain.LoginPass
		rows.Scan(&data.ID, &data.UserID, &data.Login, &data.Password, &data.MetaInfo)

		dataOutput = append(dataOutput, data)
	}

	if err = rows.Err(); err != nil {
		l.log.Err(err).Msgf("%s rows scan error", fInfo)
	}

	return dataOutput, nil
}
