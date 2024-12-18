package models

import "time"

type CRMIntegration struct {
    ID           string    `json:"id" db:"id"`
    UserID       string    `json:"user_id" db:"user_id"`
    CRMType      string    `json:"crm_type" db:"crm_type"`
    CRMAccountID string    `json:"crm_account_id" db:"crm_account_id"`
    SyncStatus   string    `json:"sync_status" db:"sync_status"`
    LastSyncedAt time.Time `json:"last_synced_at" db:"last_synced_at"`
}

