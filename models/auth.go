package models

import "github.com/gofrs/uuid"

type ValidateRoles struct {
	Role string `json:"name_role"`
}

type LogPwd struct {
	Login    string `json:"login" db:"name_account"`
	Password string `json:"password" db:"password"`
}

type Accesses struct {
	Id    uuid.UUID `json:"id" db:"id_site_role"`
	Route string    `json:"routes" db:"name_route"`
}

type AccessesId struct {
	Id uuid.UUID `json:"id" db:"id_site_role"`
}

type User struct {
	Student Student       `json:"worker"`
	Role    []RoleStudent `json:"role"`
}

func MakeUser(w Student, r []RoleStudent) User {
	return User{Student: w, Role: r}
}

type FullUser struct {
	User User   `json:"user"`
	Rt   string `json:"refresh"`
	Ac   string `json:"access"`
}

func MakeFullUser(w User, at string, rt string) FullUser {
	return FullUser{w, rt, at}
}
