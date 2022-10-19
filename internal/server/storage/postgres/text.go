package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/rs/zerolog"
)

const (
	textCreate = "textPostgres Create repo"
	textUpdate = "textPostgres Update repo"
	textDelete = "textPostgres Delete repo"
	textGetAll = "textPostgres GetAll repo"
)

// TextPostgres text data db
type TextPostgres struct {
	db  *sql.DB
	ver *Version
	log *zerolog.Logger
}

func NewTextPostgres(db *sql.DB, ver *Version, log *zerolog.Logger) *TextPostgres {
	return &TextPostgres{
		db:  db,
		ver: ver,
		log: log,
	}
}

// Create creates a text data and updates data version
func (t *TextPostgres) Create(ctx context.Context, text *domain.Text) (err error) {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		t.log.Err(err).Msgf("%s start tx error", textCreate)
		return err
	}

	defer func() {
		if err != nil {
			txError := tx.Rollback()
			if txError != nil {
				t.log.Err(txError).Msgf("%s rollback error", textCreate)
				err = fmt.Errorf("%s defer rollback error %s: %s", textCreate, txError.Error(), err.Error())
			}
		}
	}()

	row := tx.QueryRowContext(ctx,
		"INSERT INTO public.text_data(data, meta_info,user_id) VALUES ($1,$2,$3) RETURNING id", text.Text, text.MetaInfo, text.UserID)

	var id sql.NullInt64
	if err = row.Scan(&id); err != nil {
		t.log.Err(err).Msgf("%s scan id error", textCreate)
		return err
	}

	if !id.Valid {
		err = fmt.Errorf("id text data no valid")
		t.log.Err(err).Msgf("%s error", textCreate)
		return err
	}

	if err = t.ver.UpdateVersion(ctx, text.UserID, tx); err != nil {
		t.log.Err(err).Msgf("%s error", textCreate)
		return err
	}

	if err = tx.Commit(); err != nil {
		t.log.Err(err).Msgf("%s commit error", textCreate)
		return err
	}

	text.ID = int(id.Int64)

	return nil
}

// Update updates a text data and data version
func (t *TextPostgres) Update(ctx context.Context, text *domain.Text) (err error) {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		t.log.Err(err).Msgf("%s start tx error", textUpdate)
		return err
	}

	defer func() {
		if err != nil {
			txError := tx.Rollback()
			if txError != nil {
				t.log.Err(txError).Msgf("%s rollback error", textUpdate)
				err = fmt.Errorf("%s defer rollback error %s: %s", textUpdate, txError.Error(), err.Error())
			}
		}
	}()

	_, err = tx.ExecContext(ctx,
		"UPDATE public.text_data SET data=$1, meta_info=$2 WHERE id=$3 AND user_id=$4",
		text.Text, text.MetaInfo, text.ID, text.UserID)

	if err != nil {
		t.log.Err(err).Msgf("%s error", textUpdate)
		return err
	}

	if err = t.ver.UpdateVersion(ctx, text.UserID, tx); err != nil {
		t.log.Err(err).Msgf("%s error", textUpdate)
		return err
	}

	if err = tx.Commit(); err != nil {
		t.log.Err(err).Msgf("%s commit error", textUpdate)
		return err
	}

	return nil
}

// Delete deletes a text data and updates data version
func (t *TextPostgres) Delete(ctx context.Context, text *domain.Text) (err error) {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		t.log.Err(err).Msgf("%s start tx error", textDelete)
		return err
	}

	defer func() {
		if err != nil {
			txError := tx.Rollback()
			if txError != nil {
				t.log.Err(txError).Msgf("%s rollback error", textDelete)
				err = fmt.Errorf("%s defer rollback error %s: %s", textDelete, txError.Error(), err.Error())
			}
		}
	}()

	_, err = tx.ExecContext(ctx, "DELETE FROM public.text_data WHERE id=$1 AND user_id=$2", text.ID, text.UserID)
	if err != nil {
		t.log.Err(err).Msgf("%s error", textDelete)
		return err
	}

	if err = t.ver.UpdateVersion(ctx, text.UserID, tx); err != nil {
		t.log.Err(err).Msgf("%s error", textDelete)
		return err
	}

	if err = tx.Commit(); err != nil {
		t.log.Err(err).Msgf("%s commit error", textDelete)
		return err
	}

	return nil
}

// GetAll gets all text data by user id
func (t *TextPostgres) GetAll(ctx context.Context, userID int) ([]domain.Text, error) {
	rows, err := t.db.QueryContext(ctx, "SELECT id, user_id, data, meta_info FROM text_data WHERE user_id=$1", userID)
	if err != nil {
		t.log.Err(err).Msgf("%s get rows error", textGetAll)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			t.log.Err(err).Msgf("%s rows close error", textGetAll)
		}
	}()

	dataOutput := make([]domain.Text, 0)
	for rows.Next() {
		var data domain.Text
		rows.Scan(&data.ID, &data.UserID, &data.Text, &data.MetaInfo)

		dataOutput = append(dataOutput, data)
	}

	if err = rows.Err(); err != nil {
		t.log.Err(err).Msgf("%s rows scan error", textGetAll)
	}

	return dataOutput, nil
}
