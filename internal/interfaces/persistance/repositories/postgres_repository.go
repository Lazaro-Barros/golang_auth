package repositories

import (
	"database/sql"
	"fmt"

	"github.com/severusTI/auth_golang/internal/domain"
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

func (ur *UserRepository) CreateUser(user *domain.User) (err error) {
	query := "INSERT INTO users (id, name, email, phone_number, password) VALUES ($1, $2, $3, $4, $5)"
	result, err := ur.db.Exec(query, user.Id(), user.Name(), user.Email(), user.PhoneNumber(), user.EncriptedPassword())

	if err != nil {
		return ops.Err(err)
	}
	fmt.Println("aaaaaaaaaaaaaaaaaa")
	fmt.Println(result.RowsAffected())
	return nil
}

func (ur *UserRepository) GetUser(userID *string) (user *domain.User, err error) {
	query := `SELECT id, name, email, phone_number, password FROM users WHERE id = $1`
	row := ur.db.QueryRow(query, *userID)
	user = &domain.User{}
	var id, name, email, phoneNumber, encriptedPassword string
	if err := row.Scan(&id, &name, &email, &phoneNumber, &encriptedPassword); err != nil {
		return nil, ops.Err(err)
	}
	user.FillFromData(id, name, email, phoneNumber, encriptedPassword)
	return
}

func (ur *UserRepository) ListUsers() (users []domain.User, err error) {
	query := "SELECT id, name, email, phone_number, password FROM users"
	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, ops.Err(err)
	}
	defer rows.Close()

	users = []domain.User{}
	for rows.Next() {
		user := domain.User{}
		var id, name, email, phoneNumber, encriptedPassword string
		err := rows.Scan(&id, &name, &email, &phoneNumber, &encriptedPassword)
		if err != nil {
			return nil, ops.Err(err)
		}
		user.FillFromData(id, name, email, phoneNumber, encriptedPassword)
		users = append(users, user)
	}

	return
}

func (ur *UserRepository) UpdateUser(userID *string, user *domain.User) (err error) {
	query := "UPDATE users SET name = $1, email = $2, phone_number = $3, password = $4 WHERE id = $5"
	if _, err = ur.db.Exec(query, user.Name(), user.Email(), user.PhoneNumber(), user.EncriptedPassword(), *userID); err != nil {
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

// Implementação da interface Transaction para o PostgreSQL
type PostgresTransaction struct {
	tx *sql.Tx
}

func (pt *PostgresTransaction) Commit() error {
	return pt.tx.Commit()
}

func (pt *PostgresTransaction) Rollback() error {
	return pt.tx.Rollback()
}
