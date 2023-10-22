package models

type UserInfo struct {
	Fio        string `json:"FIO" db:"fio"`
	Group      string `json:"Group" db:"name_group"`
	Org        string `json:"Organization" db:"name"`
	Ebals      int    `json:"Ebals" db:"epoints"`
	Expirience int    `json:"Expirience" db:"exp"`
	Status     string `json:"Status" db:"type_role"`
}
