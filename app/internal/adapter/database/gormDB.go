package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetGormDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})

	if err != nil {
		fmt.Println("Erro ao tentar se conectar com o banco de dados.", err)
		return nil
	}
	fmt.Println("Connexão com banco de dados com sucesso!")
	return db
}