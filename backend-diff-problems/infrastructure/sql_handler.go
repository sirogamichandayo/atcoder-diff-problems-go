package infrastructure

import (
	"database/sql"
	"database/sql/driver"
	"diff-problems/config"
	"diff-problems/interfaces/database"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler(db config.SinDb) database.SqlHandler {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db.User, db.Password, db.Host, db.Port, db.Database)
	conn, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err.Error())
	}
	if err := conn.Ping(); err != nil {
		panic(err.Error())
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

func (handler *SqlHandler) Transaction(txFunc func(database.TransactionHandler) error) error {
	tx, err := handler.Conn.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()

	err = txFunc(&TransactionHandler{tx})
	return err
}

type TransactionHandler struct {
	Conn *sql.Tx
}

func (handler *TransactionHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *TransactionHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}

type Value driver.Value
