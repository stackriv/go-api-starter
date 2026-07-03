package reposiroty

import (
	"github.com/google/uuid"
	"github.com/stackriv/go-api-starter/internal/business/model"
)

type UserRepo struct {
	// db db.Database (if required)
}

func NewUserRepo( /*db *db.Database*/ ) *UserRepo {
	return &UserRepo{
		// implement database connection if required
	}
}

func (u *UserRepo) FindByEmail(email string) (*model.User, error) {
	user := new(model.User)
	// Implement database query to find user by email
	// Example: db.QueryRow("SELECT id, email, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.Password)
	return user, nil
}

func (u *UserRepo) Save(user *model.User) error {
	// Implement database query to save user
	// Example: db.Exec("INSERT INTO users (id, email, password) VALUES (?, ?, ?)", user.ID, user.Email, user.Password)
	return nil
}

func (u *UserRepo) CountUsers() (uint, error) {
	// Implement database query to count users
	// Example: var count uint; db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	return 0, nil
}

func (u *UserRepo) FindByID(id uuid.UUID) (*model.User, error) {
	user := new(model.User)
	// Implement database query to find user by ID
	// Example: db.QueryRow("SELECT id, email, password FROM users WHERE id = ?", id).Scan(&user.ID, &user.Email, &user.Password)
	return user, nil
}
