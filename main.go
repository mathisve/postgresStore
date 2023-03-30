package main

import (
	"errors"
	"log"

	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Object struct {
	ObjectName string
	Bytes      []byte
	ByteSize   int
}

func init() {
	connStr := "host=localhost port=5432 user=postgres password=password dbname=postgres sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
	}

	err = CreateSchema(db)
	if err != nil {
		log.Panic(err)
	}

}

func main() {
	// var filenames = []string{"./files/eks.png", "./files/file.txt", "./files/lofi-music.mp3"}

	// for _, file := range filenames {
	// 	data, err := os.ReadFile(file)
	// 	if err != nil {
	// 		log.Panic(err)
	// 	}

	// 	s := Object{
	// 		ObjectName: file,
	// 		Bytes:      data,
	// 		ByteSize:   len(data),
	// 	}

	// 	insertObject(db, s)
	// }

	objects, err := ListObjects(db)
	if err != nil {
		log.Println(err)
	}

	for _, obj := range objects {
		log.Println(obj)
	}

	// var c int
	// for {
	// 	for _, file := range filenames {
	// 		_, err := RetrieveObject(db, file)
	// 		if err != nil {
	// 			log.Panic(err)
	// 		}

	// 	}

	// 	c += 1
	// 	log.Println(c)
	// }

	// os.WriteFile("./files/new_eks.png", data, 0644)

}

func InsertObject(db *sql.DB, o Object) error {
	_, err := db.Exec("INSERT INTO object ( object_name, bytes, byte_size ) VALUES( $1, $2, $3)", o.ObjectName, o.Bytes, o.ByteSize)
	if err != nil {
		return err
	}

	return err
}

func RetrieveObject(db *sql.DB, objectName string) (b []byte, err error) {
	row := db.QueryRow("SELECT bytes FROM object WHERE object_name = $1", objectName)
	err = row.Scan(&b)
	if err != nil {
		return b, err
	}

	return b, err
}

func DeleteObject(db *sql.DB, objectName string) (err error) {
	res, err := db.Exec("DROP FROM object WHERE object_name = $1", objectName)
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

func ListObjects(db *sql.DB) (objects []string, err error) {
	rows, err := db.Query("SELECT object_name FROM object")
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

	return objects, err
}

func CreateSchema(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS object (
		id SERIAL PRIMARY KEY,
		object_name TEXT UNIQUE NOT NULL,
		uploaded TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		bytes BYTEA NOT NULL,
		byte_size INT NOT NULL
	);
	
	ALTER TABLE object ALTER bytes SET STORAGE EXTERNAL`)

	return err
}
