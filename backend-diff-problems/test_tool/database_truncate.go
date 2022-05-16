package test_tool

import (
	"diff-problems/config"
	"diff-problems/infrastructure"
	"diff-problems/interfaces/database"
	"os"
)

func TruncateTestTables() (database.SqlHandler, error) {
	sinDb := config.SinDb{
		Host:     os.Getenv("SIN_TEST_DB_HOST"),
		Port:     os.Getenv("SIN_TEST_DB_PORT"),
		User:     os.Getenv("SIN_TEST_DB_USER"),
		Password: os.Getenv("SIN_TEST_DB_PASSWORD"),
		Database: os.Getenv("SIN_TEST_DB_DATABASE"),
	}
	handler := infrastructure.NewSqlHandler(sinDb)

	rows, err := handler.Query("SHOW TABLES")
	if err != nil {
		return nil, err
	}
	rows.Close()

	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}
		_, err = handler.Execute("TRUNCATE " + tableName)
		if err != nil {
			return nil, err
		}
	}
	return handler, nil
}
