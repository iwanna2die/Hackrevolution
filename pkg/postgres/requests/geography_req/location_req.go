package geography_req

var GetAllLocation = `Select id_Location, code_Location,name_Location from geography.Location where deleted_at is null`

var GetLocationId = `Select id_Location, code_Location,name_Location from geography.Location where deleted_at is null and id_Location = $1`

var GetLocationIdAddress = `Select id_Location, code_Location, name_Location from geography.Location where deleted_at is null and id_Address = $1`

var GetLocationName = `Select id_Location, code_Location,name_Location from geography.Location where deleted_at is null and LOWER(name_Location) like lower('%'||$1||'%')`

var InsertLocation = `Insert into geography.Location (code_Location, name_Location,id_address) values ($1,$2,$3) returning id_Location`

var DeleteLocation = `Update geography.Location set deleted_at = now()::timestamp, updated_at = now()::timestamp where id_Location = $1`

var RestoreLocation = `Update geography.Location set deleted_at = null, updated_at = now()::timestamp where id_Location = $1`
var UpdateLocationAll = `Update geography.Location set name_Location = $2,code_location=$3,id_address=$4, updated_at = now()::timestamp where id_Location = $1`
var UpdateLocationName = `Update geography.Location set name_Location = $2, updated_at = now()::timestamp where id_Location = $1`
var UpdateLocationCode = `Update geography.Location set code_Location = $2, updated_at = now()::timestamp where id_Location = $1`
var UpdateLocationReference = `Update geography.Location set id_address = $2, updated_at = now()::timestamp where id_Location = $1`

var ShowDeleteLocation = `Select	id_Location, code_Location,name_Location from geography.Location
                                where (deleted_at +  $1) > now()::timestamp`

var ShowDeleteLocationId = `Select id_Location, code_Location,name_Location from geography.Location
                                where (deleted_at +  $1) > now()::timestamp and id_Location = $1`

var ExistsLocationName = `SELECT EXISTS(Select * from geography.Location where lower(name_Location) like lower($1))`
var ExistsLocationId = `SELECT EXISTS(Select * from geography.Location where id_location = $1)`
var ExistsLocationCode = `SELECT EXISTS(Select * from geography.Location where lower(code_location) like lower($1))`
