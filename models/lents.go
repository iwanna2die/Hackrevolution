package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type LentaUser struct {
	IdCourse    uuid.UUID `json:"IdCourse" db:"id_course"`
	File        string    `json:"File,omitempty" db:"file_path"`
	IdType      bool      `json:"IdType" db:"type_course"`
	NameCourse  string    `json:"Name_Course,omitempty" db:"name_of_course"`
	Description string    `json:"Description,omitempty" db:"description_of_course"`
}

type Notes struct {
	Text string `json:"Text,omitempty" db:"note"`
}

type Zapis struct {
	IdZapis uuid.UUID `json:"IdZapis" db:"id_post"`
	IdTask  any       `json:"IdTask,omitempty" db:"id_task"`
	Text    string    `json:"Text" db:"text"`
	TimeC   time.Time `json:"TimeC" db:"created_at"`
	TimeU   any       `json:"TimeU" db:"updated_at"`
}

type Ocenka struct {
	Fio     string `json:"Fio" db:"fio"`
	MaxMark any    `json:"Max_Mark" db:"max_mark"`
	Mark    int    `json:"Mark" db:"mark"`
	Text    string `json:"Text" db:"text"`
}
