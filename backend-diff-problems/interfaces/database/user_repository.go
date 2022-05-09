package database

import (
	"diff-problems/domain/entity"
	"fmt"
	"github.com/google/uuid"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u entity.User) (id uint64, err error) {
	uuidObj, _ := uuid.NewUUID()
	id = uint64(uuidObj.ID())
	if _, err = repo.Execute(
		"insert into users (id, first_name, last_name) values (?, ?, ?)", id, u.FirstName, u.LastName,
	); err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindById(identifier uint64) (user entity.User, err error) {
	row, err := repo.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", identifier)
	defer func(row Row) {
		err := row.Close()
		if err != nil {
			return
		}
	}(row)

	if err != nil {
		return
	}
	if !row.Next() {
		return user, fmt.Errorf("resource not found")
	}

	if err = row.Scan(&user.Id, &user.FirstName, &user.LastName); err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindAll() (users entity.Users, err error) {
	rows, err := repo.Query("SELECT id, first_name, last_name FROM users")
	defer func(rows Row) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	if err != nil {
		return
	}
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName); err != nil {
			continue
		}
		users = append(users, user)
	}
	return
}
