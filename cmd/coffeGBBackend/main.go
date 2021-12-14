package main

import (
	//"fmt"
	"log"

	//"gorm.io/gorm"

	"github.com/cofeGB/coffeGBBackend/internal/cofe_storage/gormstore/starterstore"
	//"github.com/cofeGB/coffeGBBackend/internal/cofe_storage/gormstore/navmenustore"
)

func main() {
	//Пока что бы что то было

	_, err := starterstore.StartDB("coffeDb.db")
	if err != nil {
		log.Fatal(err)
	}

}
