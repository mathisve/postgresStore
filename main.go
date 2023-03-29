package main

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

const PART_SIZE = 1000
const FILENAME = "/Users/mathis/Downloads/AIS_2021_12_15.zip"

var db *pg.DB

type Part struct {
	Filename string
	PartNum  int
	Data     []byte
}

func init() {
	db = pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "admin",
	})

	err := createSchema(db)
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	data, err := os.ReadFile(FILENAME)
	if err != nil {
		log.Panic(err)
	}

	dataLen := len(data)

	parts := dataLen / PART_SIZE

	for i := 0; i < parts; i++ {
		log.Println(len(data[PART_SIZE*i : PART_SIZE*(i+1)]))

		part := Part{
			Filename: FILENAME,
			PartNum:  i,
			Data:     data[PART_SIZE*i : PART_SIZE*(i+1)],
		}

		insertInDB(db, &part)
	}

	log.Println(len(data[PART_SIZE*parts : dataLen]))
	part := Part{

		Filename: FILENAME,
		PartNum:  parts,
		Data:     data[PART_SIZE*parts : dataLen],
	}

	insertInDB(db, &part)
}

func insertInDB(db *pg.DB, part *Part) {
	result, err := db.Model(part).Insert()
	if err != nil {
		log.Println(err)
	}

	log.Println(result.RowsAffected())

}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*Part)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
