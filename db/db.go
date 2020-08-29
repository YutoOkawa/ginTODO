package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

func InitDB() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate((&Todo{}))
	defer db.Close()
}

func InsertDB(text string, status string) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

func UpdateDB(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

func DeleteDB(id int) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

func GetAllDB() []Todo {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

func GetOneDB(id int) Todo {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
