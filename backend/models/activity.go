package models

import "time"

type Activity struct {
    ID        string    `json:"id" db:"id"`
    StepID    string    `json:"step_id" db:"step_id"`
    LeadID    string    `json:"lead_id" db:"lead_id"`
    Type      string    `json:"type" db:"type"`
    Date      time.Time `json:"date" db:"date"`
    Outcome   string    `json:"outcome" db:"outcome"`
    Notes     string    `json:"notes" db:"notes"`
}

