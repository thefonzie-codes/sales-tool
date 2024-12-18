package sqlite_connection

import (
    "database/sql"
    "log"
    "time"

    "github.com/google/uuid"
    "github.com/thefonzie-codes/sales-tracker/backend/models" // Change this import path accordingly
)

// CreateLead adds a new lead to the database
func CreateLead(lead *models.Lead) error {
    lead.ID = uuid.New().String() // Generate a unique ID for the lead
    lead.CreatedAt = time.Now()
    lead.UpdatedAt = time.Now()

    query := `INSERT INTO leads (id, name, email, phone, status, source, assigned_to, created_at, updated_at) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

    _, err := GetDB().Exec(query, lead.ID, lead.Name, lead.Email, lead.Phone, lead.Status, lead.Source, lead.AssignedTo, lead.CreatedAt, lead.UpdatedAt)
    if err != nil {
        log.Printf("Error creating lead: %v", err)
        return err
    }
    return nil
}

// GetLeadByID retrieves a lead by their ID
func GetLeadByID(id string) (*models.Lead, error) {
    query := `SELECT id, name, email, phone, status, source, assigned_to, created_at, updated_at FROM leads WHERE id = ?`
    row := GetDB().QueryRow(query, id)

    var lead models.Lead
    err := row.Scan(&lead.ID, &lead.Name, &lead.Email, &lead.Phone, &lead.Status, &lead.Source, &lead.AssignedTo, &lead.CreatedAt, &lead.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil // No lead found
        }
        log.Printf("Error retrieving lead: %v", err)
        return nil, err
    }

    return &lead, nil
}

// UpdateLead updates an existing lead's details
func UpdateLead(lead *models.Lead) error {
    lead.UpdatedAt = time.Now()

    query := `UPDATE leads SET name = ?, email = ?, phone = ?, status = ?, source = ?, assigned_to = ?, updated_at = ? WHERE id = ?`
    _, err := GetDB().Exec(query, lead.Name, lead.Email, lead.Phone, lead.Status, lead.Source, lead.AssignedTo, lead.UpdatedAt, lead.ID)
    if err != nil {
        log.Printf("Error updating lead: %v", err)
        return err
    }

    return nil
}

// DeleteLead removes a lead from the database
func DeleteLead(id string) error {
    query := `DELETE FROM leads WHERE id = ?`
    _, err := GetDB().Exec(query, id)
    if err != nil {
        log.Printf("Error deleting lead: %v", err)
        return err
    }
    return nil
}

