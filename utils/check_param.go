package utils

import (
	"HackRevolution/pkg/postgres"
	"database/sql"
	"github.com/gofrs/uuid"
	"strconv"
)

func CheckUUIDParameter(idr string) (uuid.UUID, string) {
	id, err := uuid.FromString(idr)
	if err != nil {

		e := "received invalid id path param which is not uuidV4"
		return uuid.Nil, e
	}
	return id, ""
}

func CheckIntParameter(idr string) (int, string) {
	id, err := strconv.Atoi(idr)
	if err != nil {
		e := "received invalid id path param which is not integer"
		return 0, e
	}
	return id, ""
}

func CheckExist(q string, params ...any) bool {
	rows := postgres.DB().SDB.QueryRow(q, params...)
	exist := false
	err := rows.Scan(&exist)
	switch err {
	case sql.ErrNoRows:
		return false
	case nil:
		return exist
	default:
		return false
	}
}
