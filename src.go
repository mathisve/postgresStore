package postgresStore

import (
	"context"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

type Object struct {
	ObjectName string
	Bytes      []byte
}

func (c Connection) UploadObject(o Object) (err error) {
	ctx := context.Background()
	defer ctx.Done()

	_, err = c.db.ExecContext(ctx, "INSERT INTO object ( object_name, bytes, byte_size ) VALUES( $1, $2, $3)", o.ObjectName, o.Bytes, len(o.Bytes))

	return err

}

func (c Connection) DownloadObject(objectName string) (b []byte, err error) {
	ctx := context.Background()
	defer ctx.Done()

	row := c.db.QueryRowContext(ctx, "SELECT bytes FROM object WHERE object_name = $1", objectName)
	err = row.Scan(&b)

	return b, err
}

func (c Connection) DeleteObject(objectName string) (err error) {
	ctx := context.Background()
	defer ctx.Done()

	_, err = c.db.ExecContext(ctx, "DROP FROM object WHERE object_name = $1", objectName)
	return err
}

func (c Connection) ListObjects() (objects []string, err error) {
	ctx := context.Background()
	defer ctx.Done()

	rows, err := c.db.QueryContext(ctx, "SELECT object_name FROM object")
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
