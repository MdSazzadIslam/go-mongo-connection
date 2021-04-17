package main

import (
	"fmt"
	"log"

	"github.com/MdSazzadIslam/gowithmongo/config"
)

func main() {
	fmt.Println("Server is running")
	mongoConnection, errorMongoConn := config.DBConnect()

	if errorMongoConn != nil {
		log.Println("Error when connecting to database : ", errorMongoConn.Error())
	}

	log.Println(mongoConnection)

}
