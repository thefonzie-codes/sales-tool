 package sqlite_connection

import (
    "database/sql"
    "log"
    "time"

    "github.com/google/uuid"
    "github.com/yourproject/models" // Change this import path accordingly
)

// CreateUser adds a new user to the database
func CreateUser(user *models.User) error {
    user.ID = uuid.New().String() // Generate a unique ID for the user
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()

    query := `INSERT INTO users (id, name, email, password, role, created_at, updated_at) 
              VALUES (?, ?, ?, ?, ?, ?, ?)`

    _, err := GetDB().Exec(query, user.ID, user.Name, user.Email, user.Password, user.Role, user.CreatedAt, user.UpdatedAt)
    if err != nil {
        log.Printf("Error creating user: %v", err)
        return err
    }
    return nil
}

// GetUserByID retrieves a user by their ID
func GetUserByID(id string) (*models.User, error) {
    query := `SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE id = ?`
    row := GetDB().QueryRow(query, id)

    var user models.User
    err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil // No user found
        }
        log.Printf("Error retrieving user: %v", err)
        return nil, err
    }

    return &user, nil
}

// UpdateUser updates an existing user's details
func UpdateUser(user *models.User) error {
    user.UpdatedAt = time.Now()

    query := `UPDATE users SET name = ?, email = ?, password = ?, role = ?, updated_at = ? WHERE id = ?`
    _, err := GetDB().Exec(query, user.Name, user.Email, user.Password, user.Role, user.UpdatedAt, user.ID)
    if err != nil {
        log.Printf("Error updating user: %v", err)
        return err
    }

    return nil
}

// DeleteUser removes a user from the database
func DeleteUser(id string) error {
    query := `DELETE FROM users WHERE id = ?`
    _, err := GetDB().Exec(query, id)
    if err != nil {
        log.Printf("Error deleting user: %v", err)
        return err
    }
    return nil
}

