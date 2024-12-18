package models

import "time"

type Lead struct {
    ID        string    `json:"id" db:"id"`
    Name      string    `json:"name" db:"name"`
    Email     string    `json:"email" db:"email"`
    Phone     string    `json:"phone" db:"phone"`
    Status    string    `json:"status" db:"status"`
    Source    string    `json:"source" db:"source"`
    AssignedTo string   `json:"assigned_to" db:"assigned_to"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

