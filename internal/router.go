package internal

import "net/http"

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", HealthHandler)

	

	mux.HandleFunc("/users", UsersHandler)

	mux.HandleFunc("/user/create", CreateUserHandler)

	// mux.HandleFunc("user/update", UpdateUserHandler)

	// mux.HandleFunc("user/delete", DeleteUserHandler)
	mux.HandleFunc("/", HomepageHandler)


	return mux
}
