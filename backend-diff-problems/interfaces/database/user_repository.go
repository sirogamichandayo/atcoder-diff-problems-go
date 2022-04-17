package database

import (
	"diff-problems/domain"
	"fmt"
	"github.com/google/uuid"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id uint64, err error) {
	uuidObj, _ := uuid.NewUUID()
	id = uint64(uuidObj.ID())
	fmt.Println(id)
	if _, err = repo.Execute(
		"insert into users (id, first_name, last_name) values (?, ?, ?)", id, u.FirstName, u.LastName,
	); err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindById(identifier uint64) (user domain.User, err error) {
	row, err := repo.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", identifier)
	defer row.Close()
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

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Query("SELECT id, first_name, last_name FROM users")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName); err != nil {
			continue
		}
		users = append(users, user)
	}
	return
}
