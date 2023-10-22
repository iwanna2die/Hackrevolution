package models

import "github.com/gofrs/uuid"

type Student struct {
	FullName  string    `json:"FIO" db:"fio"`
	Org       string    `json:"organization" db:"name"`
	Group     string    `json:"group,omitempty" db:"name_group"`
	IdStudent uuid.UUID `json:"id_student" db:"id_user"`
}

type RoleStudent struct {
	NameTypeStudent string `json:"name_role" db:"type_role"`
	Id              int    `json:"id_role" db:"id_role"`
}
