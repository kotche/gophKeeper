package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kotche/gophKeeper/internal/server/domain/dataType"
	"github.com/rs/zerolog"
)

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

func (l *LoginPassPostgres) Create(ctx context.Context, lp *dataType.LoginPass) (err error) {
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

	row := tx.QueryRowContext(ctx, "INSERT INTO public.login_pass(login, password, meta_info,user_id) VALUES ($1,$2,$3,$4) RETURNING id", lp.Login, lp.Password, lp.MetaInfo, lp.UserID)

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
