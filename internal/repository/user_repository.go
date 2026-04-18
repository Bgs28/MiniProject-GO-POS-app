package repository

import (
	"database/sql"
	"go-pos-app/internal/model"
	"log"
)

type UserRepository struct {
	DB *sql.DB
}


// display all user
func (r *UserRepository) GetUsers() ([]model.User, error) {
	rows, err := r.DB.Query("SELECT id, name, username, password, role FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var u model.User

		err := rows.Scan(
			&u.ID,  
			&u.Name, 
			&u.Username,
			&u.Password,
			&u.Role,
		)

		if err != nil{
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

// create user
func (r *UserRepository) CreateUsers(user model.User) error {
	query := "INSERT INTO users (name, username, password, role) VALUES (?,?,?,?)"

	_, err := r.DB.Exec(query,
	user.Name,
	user.Username,
	user.Password,
	user.Role,
	)
	
	return err
}

// update user
func (r *UserRepository) UpdateUser(user model.User) error {
	query := "UPDATE users SET name=?, username=?, password=?, role=? WHERE id=?"

	_, err := r.DB.Exec(
		query, 
		user.Name,
		user.Username,
		user.Password,
		user.Role,
		user.ID,
	)
	log.Println("Query update user id:", user.ID)

	return err
}

// delete user
func (r *UserRepository) DeleteUser(id int) error{
	query := "DELETE FROM  users WHERE id =?"

	_, err := r.DB.Exec(query, id)

	return err
}

// login
func (r *UserRepository) GetUserByUsername(username string) (model.User, error) {
	query := `
	SELECT id, username, password 
	FROM users
	WHERE username = ?
	`

	var user model.User

	err := r.DB.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
	)

	return user, err
}