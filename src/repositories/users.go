package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"

)

type Users struct {
	db *sql.DB
}

func NewRepositoryUsers(db *sql.DB) *Users {
	return &Users{db}
}

func (u Users) Create(user models.User) (uint64, error) {
	statement, err := u.db.Prepare("INSERT INTO rede_social(name, nick, email, password) VALUES(?,?,?,?)")
	if err != nil {
		return 0, nil
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, nil
	}

	lastIDIsert, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(lastIDIsert), nil
}

func (repository Users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	linhas, err := repository.db.Query("SELECT id, name, nick, email, createdIn FROM rede_social WHERE name LIKE ? OR nick LIKE ?", nameOrNick, nameOrNick)

	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var users []models.User

	for linhas.Next() {
		var user models.User

		if err = linhas.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedIn); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (repository *Users) SearchForID(ID uint64) (models.User, error) {
	linhas, err := repository.db.Query("SELECT id, name, nick, email, createdIn from rede_social where id = ?", ID)

	if err != nil {
		return models.User{}, err
	}
	defer linhas.Close()

	var user models.User

	if linhas.Next() {
		if err = linhas.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedIn); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository Users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare("UPDATE rede_social SET name = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

func (repository Users) DeleteUser(ID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM rede_social WHERE id = ?",)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repository Users) SearchForEmail(email string) (models.User, error) {
	linha, err := repository.db.Query("SELECT id, password FROM rede_social WHERE email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer linha.Close()

	var user models.User

	if linha.Next() {
		if err = linha.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}