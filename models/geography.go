package models

import "github.com/gofrs/uuid"

type CountryInp struct {
	CodeCountry string ` db:"code_country" json:"code_country"`
	NameCountry string ` db:"name_country" json:"name_country"`
}

type CountryUpd struct {
	CodeCountry string ` db:"code_country" json:"code_country,omitempty"`
	NameCountry string ` db:"name_country" json:"name_country,omitempty"`
}

type RegionInp struct {
	CodeRegion string `json:"code_region"  db:"code_region"`
	NameRegion string `json:"name_region"  db:"name_region"`
	IdCountry  int    `json:"id_country"`
}

type RegionUpd struct {
	CodeRegion string `json:"code_region,omitempty"  db:"code_region"`
	NameRegion string `json:"name_region,omitempty"  db:"name_region"`
	IdCountry  int    `json:"id_country,omitempty" db:"id_country"`
}

type CityInp struct {
	CodeCity string `json:"code_city"  db:"code_city"`
	NameCity string `json:"name_city"  db:"name_city"`
	IdRegion int    `json:"id_region" db:"id_region"`
}

type CityUpd struct {
	CodeCity string `json:"code_city,omitempty"  db:"code_city"`
	NameCity string `json:"name_city,omitempty"  db:"name_city"`
	IdRegion int    `json:"id_region,omitempty" db:"id_region"`
}

type AddressInp struct {
	PostCode    string `json:"post_code" db:"postcode"`
	CodeAddress string `json:"code_address"`
	NameAddress string `json:"name_address"`
	IdCity      int    `json:"id_city"`
}

type AddressUpd struct {
	PostCode    string `json:"post_code,omitempty" db:"postcode"`
	CodeAddress string `json:"code_address,omitempty"  db:"code_address"`
	NameAddress string `json:"name_address,omitempty"  db:"name_address"`
	IdCity      int    `json:"id_city,omitempty" db:"id_city"`
}

type LocationInp struct {
	CodeLocation string    `json:"code_location" db:"code_location"`
	NameLocation string    `json:"name_location" db:"name_location"`
	IdAddress    uuid.UUID `json:"id_address" db:"id_address"`
}

type LocationUpd struct {
	CodeLocation string    `json:"code_location,omitempty" db:"code_location"`
	NameLocation string    `json:"name_location,omitempty" db:"name_location"`
	IdAddress    uuid.UUID `json:"id_address,omitempty" db:"id_address"`
}

type WorkspaceInp struct {
	CodeWorkspace string    `json:"code_workspace"  db:"code_workspace"`
	NameWorkspace string    `json:"name_workspace"  db:"name_workspace"`
	IdLocation    uuid.UUID `json:"id_location"  db:"id_workspace"`
}

type WorkspaceUpd struct {
	CodeWorkspace string    `json:"code_workspace,omitempty"  db:"code_workspace"`
	NameWorkspace string    `json:"name_workspace,omitempty"  db:"name_workspace"`
	IdLocation    uuid.UUID `json:"id_location,omitempty"  db:"id_workspace"`
}

type Country struct {
	IdCountry   int    ` db:"id_country" json:"id_country"`
	CodeCountry string ` db:"code_country" json:"code_country"`
	NameCountry string ` db:"name_country" json:"name_country"`
}

type Region struct {
	IdRegion   int    `json:"id_region"  db:"id_region"`
	CodeRegion string `json:"code_region"  db:"code_region"`
	NameRegion string `json:"name_region"  db:"name_region"`
}

type City struct {
	IdCity   int    `json:"id_city"  db:"id_city"`
	CodeCity string `json:"code_city"  db:"code_city"`
	NameCity string `json:"name_city"  db:"name_city"`
}

type Address struct {
	IdAddress   uuid.UUID `json:"id_address"  db:"id_address"`
	CodeAddress string    `json:"code_address"  db:"code_address"`
	NameAddress string    `json:"name_address"  db:"name_address"`
	PostCode    string    `json:"post_code" db:"postcode"`
}

type Location struct {
	IdLocation   uuid.UUID `json:"id_location"  db:"id_location"`
	CodeLocation string    `json:"code_location"  db:"code_location"`
	NameLocation string    `json:"name_location"  db:"name_location"`
}

type Workspace struct {
	IdWorkspace   uuid.UUID `json:"id_workspace"  db:"id_workspace"`
	CodeWorkspace string    `json:"code_workspace"  db:"code_workspace"`
	NameWorkspace string    `json:"name_workspace"  db:"name_workspace"`
}

type CountryTrace struct {
	Country     `json:"country"`
	RegionTrace //`json:"region"`
}

type RegionTrace struct {
	Region    `json:"region"`
	CityTrace //`json:"city"`
}

type CityTrace struct {
	City         `json:"city"`
	AddressTrace //`json:"location"`
}

type AddressTrace struct {
	Address       `json:"address"`
	LocationTrace //`json:"location"`
}

type LocationTrace struct {
	Location  `json:"location"`
	Workspace `json:"workspace"`
}

//func makeTrace(c Country, r Region, ct City, a Address, l Location, w Workspace) CountryTrace {
//	return CountryTrace{Co: c, Reg: RegionTrace{Reg: r, Ct: CityTrace{Ct: ct, Add: AddressTrace{Add: a, Loc: LocationTrace{Loc: l, Ws: w}}}}}
//}
