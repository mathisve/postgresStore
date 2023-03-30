package postgresStore

import (
	"log"

	_ "github.com/lib/pq"
)

func init() {
	log.Println("postgres-store running ...")
}
