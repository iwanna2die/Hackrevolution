package models

import "github.com/gofrs/uuid"

type LogsInp struct {
	IdUser      uuid.UUID `json:"user id" db:"id_user"`
	IdSubject   string    `json:"subject id" db:"id_subject"`
	IdAction    string    `json:"id of action" db:"id_action"`
	ValueBefore string    `json:"value before,omitempty" db:"before"`
	ValueAfter  string    `json:"value after,omitempty" db:"after"`
	Description string    `json:"description,omitempty" db:"description"`
	Success     int       `json:"is operation made" db:"success"`
}
