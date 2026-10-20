package models

import "time"

type AuditLog struct {
	ID        uint        `json:"id"`
	UserID    uint        `json:"user_id"`
	Action    string      `json:"action"`
	Table     string      `json:"table"`
	OldValue  interface{} `json:"old_value,omitempty"`
	NewValue  interface{} `json:"new_value,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}
