package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kotche/gophKeeper/internal/server/domain"
	"github.com/rs/zerolog"
)

// BankCardPostgres bank card data db
type BankCardPostgres struct {
	db  *sql.DB
	log *zerolog.Logger
}

func NewBankCardPostgres(db *sql.DB, log *zerolog.Logger) *BankCardPostgres {
	return &BankCardPostgres{
		db:  db,
		log: log,
	}
}

// Create creates a bank card data and updates data version
func (b *BankCardPostgres) Create(ctx context.Context, bank *domain.BankCard) (err error) {
	const fInfo = "bankCardPostgres create repo"

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
		"INSERT INTO public.bank_card(number, meta_info,user_id) VALUES ($1,$2,$3) RETURNING id", bank.Number, bank.MetaInfo, bank.UserID)

	var id sql.NullInt64
	if err = row.Scan(&id); err != nil {
		b.log.Err(err).Msgf("%s scan id error", fInfo)
		return err
	}

	if !id.Valid {
		err = fmt.Errorf("id bank card data no valid")
		b.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = updateVersion(ctx, bank.UserID, tx, b.log); err != nil {
		b.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = tx.Commit(); err != nil {
		b.log.Err(err).Msgf("%s commit error", fInfo)
		return err
	}

	bank.ID = int(id.Int64)

	return nil
}

// Update updates a bank card data and data version
func (b *BankCardPostgres) Update(ctx context.Context, bank *domain.BankCard) (err error) {
	const fInfo = "bankCardPostgres update repo"

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
		"UPDATE public.bank_card SET number=$1,meta_info=$2 WHERE id=$3 AND user_id=$4",
		bank.Number, bank.MetaInfo, bank.ID, bank.UserID)

	if err != nil {
		b.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = updateVersion(ctx, bank.UserID, tx, b.log); err != nil {
		b.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = tx.Commit(); err != nil {
		b.log.Err(err).Msgf("%s commit error", fInfo)
		return err
	}

	return nil
}

// Delete deletes a bank card data and updates data version
func (b *BankCardPostgres) Delete(ctx context.Context, bank *domain.BankCard) (err error) {
	const fInfo = "bankCardPostgres delete repo"

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

	_, err = tx.ExecContext(ctx, "DELETE FROM public.bank_card WHERE id=$1 AND user_id=$2", bank.ID, bank.UserID)
	if err != nil {
		b.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = updateVersion(ctx, bank.UserID, tx, b.log); err != nil {
		b.log.Err(err).Msgf("%s error", fInfo)
		return err
	}

	if err = tx.Commit(); err != nil {
		b.log.Err(err).Msgf("%s commit error", fInfo)
		return err
	}

	return nil
}

// GetAll returns all bank card data by user id
func (b *BankCardPostgres) GetAll(ctx context.Context, userID int) ([]domain.BankCard, error) {
	const fInfo = "BankCardPostgres getAll repo"

	rows, err := b.db.QueryContext(ctx, "SELECT id, user_id, number, meta_info FROM bank_card WHERE user_id=$1", userID)
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

	dataOutput := make([]domain.BankCard, 0)
	for rows.Next() {
		var data domain.BankCard
		rows.Scan(&data.ID, &data.UserID, &data.Number, &data.MetaInfo)

		dataOutput = append(dataOutput, data)
	}

	if err = rows.Err(); err != nil {
		b.log.Err(err).Msgf("%s rows scan error", fInfo)
	}

	return dataOutput, nil
}
