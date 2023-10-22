package models

import "github.com/gofrs/uuid"

type TypeWorkplaceInp struct {
	NameTypeWp string `json:"name_type_workplace" db:"name_type_workplace"`
}

type TypeWorkplace struct {
	IdTypeWp   string `json:"id_type_workplace" db:"id_type_workplace"`
	NameTypeWp string `json:"name_type_workplace" db:"name_type_workplace"`
}

type WorkplaceInp struct {
	Description string    `json:"description" db:"description"`
	IdTypeWp    uuid.UUID `json:"id_type_workplace" db:"id_type_workplace"`
	IdWorkspace uuid.UUID `json:"id_workspace" db:"id_workspace"`
}

type Workplace struct {
	IdWorkplace uuid.UUID `json:"id_workplace" db:"id_workplace"`
	Description string    `json:"description" db:"description"`
	IdTypeWp    uuid.UUID `json:"id_type_workplace" db:"id_type_workplace"`
	IdWorkspace uuid.UUID `json:"id_workspace" db:"id_workspace"`
}
