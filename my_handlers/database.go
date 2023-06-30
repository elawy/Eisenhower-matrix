package my_handlers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

func createTable(db *gorm.DB, tableName string) {
	table := db.Table(tableName)
	if !table.Migrator().HasTable(tableName) {
		table.AutoMigrate(&TodoItem{})
	}
}

func ToCreate(table_1 string, table_2 string, table_3 string, table_4 string) {

	createTable(db, table_1)
	createTable(db, table_3)
	createTable(db, table_2)
	createTable(db, table_4)

}
