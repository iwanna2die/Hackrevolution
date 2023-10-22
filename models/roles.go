package models

import "github.com/gofrs/uuid"

type SiteRouteInp struct {
	NameSiteRoute string `json:"name_site_route" db:"name_route"`
	Description   string `json:"description,omitempty" db:"description_role"`
}

type SiteRouteDesc struct {
	Description string `json:"description,omitempty" db:"description_role"`
}

type SiteRoute struct {
	Id            uuid.UUID `json:"id" db:"id_site_role"`
	NameSiteRoute string    `json:"name_site_route" db:"name_route"`
	Description   any       `json:"description,omitempty" db:"description_role"`
}

type UserRoleInp struct {
	NameUserRole string `json:"name_role" db:"name_type_worker"`
}

type UserRole struct {
	IdRole       uuid.UUID `json:"id_role" db:"id_type_worker"`
	NameUserRole string    `json:"name_role" db:"name_type_worker"`
}

type RouteIDs struct {
	Id uuid.UUID `json:"id" db:"id_site_role"`
}

type ConnectionRoleInp struct {
	Routes []RouteIDs `json:"routes"`
}

type FullRoleInp struct {
	Role   RoleStudent `json:"role_worker"`
	Routes []Accesses  `json:"routes"`
}
