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
		return "Error creating user"
	}
	

	
	return "User created successfully"
}

func UpdateUser(r *http.Request) string {

	err := r.ParseForm()
	if err != nil {
		return "Invalid form data"
	}
	id := r.FormValue("id")
	name := r.FormValue("name")
	email := r.FormValue("email")
	query := `UPDATE users SET name=$1, email=$2 WHERE id=$3`

	if id == "" || name == "" || email == "" {
		return "ID, Name and Email are required"
	}

	 _ ,err = DB.Exec(query, name, email, id)
	 
	if err != nil {
		log.Println("Error updating user:", err)
		return "Error updating user"
	}
	return "User updated successfully"
}

func DeleteUser(r *http.Request) string {

	err := r.ParseForm()
	if err != nil {
		return "Invalid form data"
	}
	id := r.FormValue("id")
	query := `DELETE FROM users WHERE id=$1`

	if id == "" {
		return "ID is required"
	}

	 _ ,err = DB.Exec(query, id)
	 
	if err != nil {
		log.Println("Error deleting user:", err)
		return "Error deleting user"
	}
	log.Println("user deleted with Id: ", id)
	return "User deleted successfully"
}
