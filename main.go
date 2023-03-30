package main

import (
	"log"
	"os"

	po "postgres_object/src"
)

func main() {
	// create new connection
	c, err := po.NewConnection(po.DefaultConnectionConfig)
	if err != nil {
		log.Panic(err)
	}

	var filenames = []string{"./etc/files/eks.png", "./etc/files/file.txt", "./etc/files/lofi-music.mp3"}

	for _, file := range filenames {
		data, err := os.ReadFile(file)
		if err != nil {
			log.Panic(err)
		}

		s := po.Object{
			ObjectName: file,
			Bytes:      data,
			ByteSize:   len(data),
		}

		err = c.InsertObject(s)
		if err != nil {
			log.Println(err)
		}
	}

	objects, err := c.ListObjects()
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
