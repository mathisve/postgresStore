package postgresStore

import (
	"context"
	"database/sql"
	"fmt"
)

type Connection struct {
	db *sql.DB
}

// NewConnection is used to create a connection to the Postgres database
// takes in ConnectionConfig object which contains the connection configuration.
// returns Connection object, contains the db object.
func NewConnection(config ConnectionConfig) (c Connection, err error) {
	var connStr string

	if config.ConnStr == "" {
		connStr = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
			config.Host,
			config.Port,
			config.Username,
			config.Password,
			config.DBName,
			config.SslMode,
		)
	} else {
		connStr = config.ConnStr
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return c, err
	}

	// leave it up to the user
	// configure this in postgresql.conf
	// max_connections param
	db.SetMaxOpenConns(0)

	// log.Println("connected to db")

	err = createSchema(db, config)
	if err != nil {
		return c, err
	}

	return Connection{
		db: db,
	}, err
}

// createSchema is used to create the correct object schema in the database
// takes in db object and storageMode string
// storageMode string is used to alter the bytes storage mode in postgres
// storageModeExtended means TOAST with compression
// storageModeExternal means TOAST no compression
// is not to be used outside the package
func createSchema(db *sql.DB, c ConnectionConfig) error {

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if c.Unlogged {
		_, err = tx.ExecContext(ctx, `CREATE UNLOGGED TABLE IF NOT EXISTS object (
			id SERIAL PRIMARY KEY,
			object_name TEXT UNIQUE NOT NULL,
			uploaded TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			bytes BYTEA NOT NULL,
			byte_size INT NOT NULL);`)
		if err != nil {
			return err
		}
	} else {
		_, err = tx.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS object (
			id SERIAL PRIMARY KEY,
			object_name TEXT UNIQUE NOT NULL,
			uploaded TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			bytes BYTEA NOT NULL,
			byte_size INT NOT NULL);`)
		if err != nil {
			return err
		}
	}

	_, err = tx.ExecContext(ctx, `CREATE UNIQUE INDEX IF NOT EXISTS object_name_idx ON object(object_name);`)
	if err != nil {
		return err
	}

	if c.StorageMode == StorageModeExtended {
		_, err = tx.ExecContext(ctx, `ALTER TABLE object ALTER bytes SET STORAGE EXTENDED`)
	} else {
		_, err = tx.ExecContext(ctx, `ALTER TABLE object ALTER bytes SET STORAGE EXTERNAL`)
	}

	if err != nil {
		return err
	}

	return tx.Commit()
}
