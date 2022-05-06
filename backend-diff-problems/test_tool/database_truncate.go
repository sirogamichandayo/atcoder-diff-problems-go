package test_tool

import (
	"diff-problems/interfaces/database"
	"log"
)

func TruncateTables(db database.SqlHandler) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatal("show tables error:", err)
	}

	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			log.Fatal("show table error:", err)
		}
		_, err = db.Execute("TRUNCATE " + tableName)
		if err != nil {
			log.Fatal("truncate table error:", err)
		}
	}
}
