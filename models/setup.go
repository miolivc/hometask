package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupModels() *gorm.DB {

	pg_conn := fmt.Sprintf("host=localhost port=5432 user=postgres dbname=hometask password=postgres sslmode=disable")

	fmt.Println("conname is\t\t", pg_conn)
	db, err := gorm.Open("postgres", pg_conn)
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Task{})

	// Initialise value
	var tasks = []Task{
		{Name: "Limpar chão da cozinha", Level: 2, Daily: false},
		{Name: "Limpar chão da varanda", Level: 2, Daily: false},
		{Name: "Limpar caixa de areia", Level: 1, Daily: true},
	}

	for _, task := range tasks {
		db.Create(&task)
	}

	return db
}
