package repository

import (
	"database/sql"
	"go-pos-app/internal/model"
)

type UserRepository struct {
	DB *sql.DB
}


// menampilkan semua user
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
	// log.Println("Query update user id:", user.ID)

	return err
}

func (r *UserRepository) DeleteUser(id int) error{
	query := "DELETE FROM  users WHERE id =?"

	_, err := r.DB.Exec(query, id)

	return err
}