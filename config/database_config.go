package config

import (
	"be-todo-app/database"
	"be-todo-app/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db_user = "root"
var db_password = ""
var db_host = "localhost"
var db_port = "3306"
var db_name = "fiber_todo_app"

func BootDatabase() {
	if dbNameEnv := os.Getenv("DB_NAME"); dbNameEnv != "" {
		db_name = dbNameEnv
	}

	if dbPortEnv := os.Getenv("DB_PORT"); dbPortEnv != "" {
		db_port = dbPortEnv
	}

	if dbUserEnv := os.Getenv("DB_USER"); dbUserEnv != "" {
		db_user = dbUserEnv
	}

	if dbPasswordEnv := os.Getenv("DB_PASSWORD"); dbPasswordEnv != "" {
		db_password = dbPasswordEnv
	}

	if dbHostEnv := os.Getenv("DB_HOST"); dbHostEnv != "" {
		db_host = dbHostEnv
	}
}

func ConnectDatabase() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password, db_host, db_port, db_name)
	database.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Tidak dapat connect ke database")
	} else {
		log.Println("Berhasil connect ke database")
	}
}

func RunMigration() {
	err := database.DB.AutoMigrate(
		model.Todo{},
	)

	if err != nil {
		fmt.Println(err)
		log.Println("Gagal membuat migration")
	} else {
		log.Println("Berhasil menjalankan migration")
	}
}
