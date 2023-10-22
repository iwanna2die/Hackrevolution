package postgres

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
	"reflect"
)

func ReadDbAll(rows *sqlx.Rows, err error, inter any) (int, any) {
	switch err {
	case sql.ErrNoRows:
		defer rows.Close()
		return http.StatusNotFound, "no rows found"
	case nil:
		defer rows.Close()
		a := make([]any, 0)

		rowsErr := false
		for rows.Next() {
			newinter := CloneValue(inter)
			e := rows.StructScan(newinter)
			fmt.Println(e)
			if e != nil {
				rowsErr = true
				break
			}
			a = append(a, newinter)
		}
		if rowsErr {
			return 500, "cant take rows"
		}
		fmt.Println(a)
		return 200, a
	default:
		defer rows.Close()
		return 500, "database error"
	}
}

func CloneValue(v any) any {
	if v == nil {
		return nil
	}
	var clonedValue any
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		clonedValue = reflect.New(val.Elem().Type()).Interface()
	} else {
		clonedValue = reflect.New(val.Type()).Elem().Interface()
	}
	return clonedValue
}

func Changes(status int, q, nameStatus string, args ...any) (int, any) {
	_, err := DB().SDB.Exec(q, args...)
	if err != nil {
		return 400, err.Error()
	}
	return status, nameStatus
}

func Create(q string, arg ...any) (int, any) {
	var id any
	err := DB().SDB.Get(&id, q, arg...)
	if err != nil {
		return 500, err.Error()
	}
	return 200, id
}
