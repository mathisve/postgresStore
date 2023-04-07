package postgresStore

import (
	"context"
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

type Object struct {
	ObjectName string
	Bytes      []byte
}

func (c Connection) UploadObject(o Object) error {
	ctx := context.Background()

	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO object ( object_name, bytes, byte_size ) VALUES( $1, $2, $3)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(o.ObjectName, o.Bytes, len(o.Bytes))

	if err != nil {
		err = tx.Rollback()
		return err
	}

	return tx.Commit()

}

func (c Connection) DownloadObject(objectName string) (b []byte, err error) {
	row := c.db.QueryRow("SELECT bytes FROM object WHERE object_name = $1", objectName)
	err = row.Scan(&b)
	if err != nil {
		return b, err
	}

	return b, err
}

func (c Connection) DeleteObject(objectName string) (err error) {
	res, err := c.db.Exec("DROP FROM object WHERE object_name = $1", objectName)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		err = errors.New("no objects deleted")
	}

	return err
}

func (c Connection) ListObjects() (objects []string, err error) {
	rows, err := c.db.Query("SELECT object_name FROM object")
	if err != nil {
		return objects, err
	}

	for rows.Next() {
		var object_name string
		err = rows.Scan(&object_name)
		if err != nil {
			log.Println(err)
		}

		objects = append(objects, object_name)
	}

	if len(objects) == 0 {
		return objects, errors.New("no objects found")
	}

	return objects, err
}
