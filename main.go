package main

import (
	"os"

	"github.com/filipegms5/cotemig-projeto-software/database"
	"github.com/filipegms5/cotemig-projeto-software/router"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Instruction os how to run the project on README
func main() {

	// Connect to the database
	db, err := gorm.Open(sqlite.Open("adoteUmIdoso.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.SetupDatabase(db)
	router := router.SetupRouter(db)

	router.Run(os.Getenv("PORT"))
}
