package main

import (
	"log"
	"os"

	"github.com/mathisve/postgresStore"
)

const (
	filename = "file.txt"
)

func main() {

	// create connection
	c, err := postgresStore.NewConnection(postgresStore.DefaultConnectionConfig.SetUnlogged(true))
	if err != nil {
		log.Println(err)
	}

	// open file
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Println(err)
	}

	log.Println(data)

	// upload objects
	err = c.UploadObject(postgresStore.Object{
		ObjectName: filename,
		Bytes:      data,
	})

	if err != nil {
		log.Println(err)
	}

	// list objects
	objects, err := c.ListObjects()
	if err != nil {
		log.Println(err)
	}

	for _, obj := range objects {
		log.Println(obj)
	}
}
