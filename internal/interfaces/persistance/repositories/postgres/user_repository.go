package postgres

import (
	"database/sql"

	"github.com/severusTI/auth_golang/internal/domain/entities"
	"github.com/severusTI/auth_golang/pkg/ops"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(user *entities.User) (err error) {
	query := "INSERT INTO users (id, name, email, phone_number, password) VALUES ($1, $2, $3, $4, $5)"
	_, err = ur.db.Exec(query, user.Id(), user.Name(), user.Email(), user.PhoneNumber(), user.HashedPassword())
	if err != nil {
		return ops.Err(err)
	}
	return nil
}

func (ur *UserRepository) GetUserByEmail(UserEmail *string) (user *entities.User, err error) {
	query := `SELECT id, name, email, phone_number, password FROM users WHERE LOWER(email) = LOWER($1)`
	row := ur.db.QueryRow(query, *UserEmail)
	user = &entities.User{}
	var id, name, email, phoneNumber, hashPassword string
	if err := row.Scan(&id, &name, &email, &phoneNumber, &hashPassword); err != nil {
		return nil, ops.Err(err)
	}
	user.FillFromData(id, name, email, phoneNumber, hashPassword)
	return
}

func (ur *UserRepository) UpdateUser(userID *string, user *entities.User) (err error) {
	query := "UPDATE users SET name = $1, email = $2, phone_number = $3, password = $4 WHERE id = $5"
	if _, err = ur.db.Exec(query, user.Name(), user.Email(), user.PhoneNumber(), user.HashedPassword(), *userID); err != nil {
		return ops.Err(err)
	}
	return
}

func (ur *UserRepository) DeleteUser(userID *string) (err error) {
	query := "DELETE FROM users WHERE id = $1"
	if _, err = ur.db.Exec(query, *userID); err != nil {
		return ops.Err(err)
	}
	return
}
