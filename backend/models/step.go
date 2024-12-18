package models

import "time"

type Step struct {
    ID        string    `json:"id" db:"id"`
    CadenceID string    `json:"cadence_id" db:"cadence_id"`
    Action    string    `json:"action" db:"action"`
    DueDate   time.Time `json:"due_date" db:"due_date"`
    Status    string    `json:"status" db:"status"`
    Notes     string    `json:"notes" db:"notes"`
}

