package models

import "github.com/gofrs/uuid"

type DataInfoInp struct {
	NameDataInfo string `json:"name_data_info" db:"name_data"`
}

type DataInfo struct {
	IdDataInfo   uuid.UUID `json:"id_data_info" db:"id_data"`
	NameDataInfo string    `json:"name_data_info" db:"name_data"`
}

type DataParamInp struct {
	NameDataParam uuid.UUID `json:"name_data_param"`
}
