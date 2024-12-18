package models

import "time"

type FollowUp struct {
    ID        string    `json:"id" db:"id"`
    ActivityID string   `json:"activity_id" db:"activity_id"`
    LeadID    string    `json:"lead_id" db:"lead_id"`
    DueDate   time.Time `json:"due_date" db:"due_date"`
    Status    string    `json:"status" db:"status"`
    Notes     string    `json:"notes" db:"notes"`
}

