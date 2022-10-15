package postgres

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// PGX postgresSQL driver
type PGX struct {
	DB *sql.DB
}

func NewPGX(DSN string) (*PGX, error) {
	db, err := sql.Open("pgx", DSN)
	if err != nil {
		return nil, err
	}
	pgx := &PGX{DB: db}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = pgx.Init()
	if err != nil {
		return nil, err
	}

	return pgx, nil
}

// Ping checks the connection to the database
func (p *PGX) Ping() error {
	if err := p.DB.Ping(); err != nil {
		return err
	}
	return nil
}

// Init creates database tables
func (p *PGX) Init() error {
	_, err := p.DB.Exec(`
			CREATE TABLE IF NOT EXISTS public.users(
		    id BIGSERIAL PRIMARY KEY,
    		login TEXT NOT NULL UNIQUE,
    		password TEXT NOT NULL);

			CREATE TABLE IF NOT EXISTS public.versions(
			    version BIGINT NOT NULL,
				user_id BIGINT NOT NULL,
				CONSTRAINT uniq_version_user_id UNIQUE (version, user_id),
				FOREIGN KEY (user_id) REFERENCES public.users (id));

			CREATE TABLE IF NOT EXISTS public.login_pass(
				id BIGSERIAL PRIMARY KEY,
				login TEXT NOT NULL,
				password TEXT NOT NULL,
				meta_info TEXT NOT NULL,
				user_id BIGINT NOT NULL,
				FOREIGN KEY (user_id) REFERENCES public.users (id));

			CREATE TABLE IF NOT EXISTS public.text_data(
				id BIGSERIAL PRIMARY KEY,
				data TEXT NOT NULL,
				meta_info TEXT NOT NULL,
				user_id BIGINT NOT NULL,
				FOREIGN KEY (user_id) REFERENCES public.users (id));

			CREATE TABLE IF NOT EXISTS public.binary_data(
				id BIGSERIAL PRIMARY KEY,
				data TEXT NOT NULL,
				meta_info TEXT NOT NULL,
				user_id BIGINT NOT NULL,
				FOREIGN KEY (user_id) REFERENCES public.users (id));

			CREATE TABLE IF NOT EXISTS public.bank_card(
				id BIGSERIAL PRIMARY KEY,
				number TEXT NOT NULL,
				meta_info TEXT NOT NULL,
				user_id BIGINT NOT NULL,
				FOREIGN KEY (user_id) REFERENCES public.users (id));
`)

	if err != nil {
		return err
	}

	return nil
}
