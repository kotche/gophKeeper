package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/rs/zerolog"
)

// BinaryPostgres binary data db
type BinaryPostgres struct {
	db  *sql.DB
	log *zerolog.Logger
}

func NewBinaryPostgres(db *sql.DB, log *zerolog.Logger) *BinaryPostgres {
	return &BinaryPostgres{
		db:  db,
		log: log,
	}
}

// Create creates a binary data and updates data version
func (b *BinaryPostgres) Create(ctx context.Context, bin *domain.Binary) (err error) {
	const fInfo = "binaryPostgres create repo"

	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		b.log.Err(err).Msgf("%s start tx error", fInfo)
		return err
	}

	defer func() {
		if err != nil {
			txError := tx.Rollback()
			if txError != nil {
				b.log.Err(txError).Msgf("%s rollback error", fInfo)
				err = fmt.Errorf("%s defer rollback error %s: %s", fInfo, txError.Error(), err.Error())
			}
		}
	}()

	row := tx.QueryRowContext(ctx,
		"INSERT INTO public.binary_data(data, meta_info,user_id) VALUES ($1,$2,$3) RETURNING id", bin.Binary, bin.MetaInfo, bin.UserID)

	var id sql.NullInt64
	if err = row.Scan(&id); err != nil {
		b.log.Err(err).Msgf("%s scan id error", fInfo)
		return err
	}

	if !id.Valid {
		err = fmt.Errorf("id binary data no valid")
		b.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = updateVersion(ctx, bin.UserID, tx, b.log); err != nil {
		b.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = tx.Commit(); err != nil {
		b.log.Err(err).Msgf("%s commit error", fInfo)
		return err
	}

	bin.ID = int(id.Int64)

	return nil
}

// Update updates a binary data and data version
func (b *BinaryPostgres) Update(ctx context.Context, bin *domain.Binary) (err error) {
	const fInfo = "binaryPostgres update repo"

	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		b.log.Err(err).Msgf("%s start tx error", fInfo)
		return err
	}

	defer func() {
		if err != nil {
			txError := tx.Rollback()
			if txError != nil {
				b.log.Err(txError).Msgf("%s rollback error", fInfo)
				err = fmt.Errorf("%s defer rollback error %s: %s", fInfo, txError.Error(), err.Error())
			}
		}
	}()

	_, err = tx.ExecContext(ctx,
		"UPDATE public.binary_data SET data=$1,meta_info=$2 WHERE id=$3 AND user_id=$4",
		bin.Binary, bin.MetaInfo, bin.ID, bin.UserID)

	if err != nil {
		b.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = updateVersion(ctx, bin.UserID, tx, b.log); err != nil {
		b.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = tx.Commit(); err != nil {
		b.log.Err(err).Msgf("%s commit error", fInfo)
		return err
	}

	return nil
}

// Delete deletes a binary data and updates data version
func (b *BinaryPostgres) Delete(ctx context.Context, bin *domain.Binary) (err error) {
	const fInfo = "BinaryPostgres delete repo"

	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		b.log.Err(err).Msgf("%s start tx error", fInfo)
		return err
	}

	defer func() {
		if err != nil {
			txError := tx.Rollback()
			if txError != nil {
				b.log.Err(txError).Msgf("%s rollback error", fInfo)
				err = fmt.Errorf("%s defer rollback error %s: %s", fInfo, txError.Error(), err.Error())
			}
		}
	}()

	_, err = tx.ExecContext(ctx, "DELETE FROM public.binary_data WHERE id=$1 AND user_id=$2", bin.ID, bin.UserID)
	if err != nil {
		b.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = updateVersion(ctx, bin.UserID, tx, b.log); err != nil {
		b.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = tx.Commit(); err != nil {
		b.log.Err(err).Msgf("%s commit error", fInfo)
		return err
	}

	return nil
}

// GetAll returns all binary data by user id
func (b *BinaryPostgres) GetAll(ctx context.Context, userID int) ([]domain.Binary, error) {
	const fInfo = "binaryPostgres getAll repo"

	rows, err := b.db.QueryContext(ctx, "SELECT id, user_id, data, meta_info FROM binary_data WHERE user_id=$1", userID)
	if err != nil {
		b.log.Err(err).Msgf("%s get rows error", fInfo)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			b.log.Err(err).Msgf("%s rows close error", fInfo)
		}
	}()

	dataOutput := make([]domain.Binary, 0)
	for rows.Next() {
		var data domain.Binary
		rows.Scan(&data.ID, &data.UserID, &data.Binary, &data.MetaInfo)

		dataOutput = append(dataOutput, data)
	}

	if err = rows.Err(); err != nil {
		b.log.Err(err).Msgf("%s rows scan error", fInfo)
	}

	return dataOutput, nil
}
