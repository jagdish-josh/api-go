package internal

import (
	"log"
	"net/http"
)

func HealthService() string {
	return "OK"
}

func GetAllUsers() ([]User, error) {
	query := `SELECT id, name, email FROM users`

	rows, err := DB.Query(query)
	
	if err != nil {
		log.Println("Error querying users:", err)
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
		
	}

	return users, nil
}

func CreateUser(r *http.Request) string{

	name := r.FormValue("name")
	email := r.FormValue("email")

	if name == "" || email == "" {
		
		return "Name and email are required"
	}

	query := `INSERT INTO users (name, email) VALUES ($1, $2)`
	_, err := DB.Exec(query, name, email)


	if err != nil {
		log.Println("Error creating user:", err)
	}
	

	
	return "User created successfully"
}


