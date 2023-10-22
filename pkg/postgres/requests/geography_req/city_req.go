package geography_req

var GetAllCity = `Select id_City, code_City,name_City from geography.City where deleted_at is null`

var GetCityId = `Select id_City, code_City,name_City from geography.City where deleted_at is null and id_City = $1`

var GetCityIdRegion = `Select id_City, code_City, name_City from geography.City where deleted_at is null and id_Region = $1`

var GetCityName = `Select id_City, code_City,name_City from geography.City where deleted_at is null and LOWER(name_City) like lower('%'||$1||'%')`

var InsertCity = `Insert into geography.City (code_City, name_City,id_region) values ($1,$2,$3) returning id_City`

var DeleteCity = `Update geography.City set deleted_at = now()::timestamp, updated_at = now()::timestamp where id_City = $1`

var RestoreCity = `Update geography.City set deleted_at = null, updated_at = now()::timestamp where id_City = $1`

var UpdateCityName = `Update geography.City set name_City = $2, updated_at = now()::timestamp where id_region = $1`

var UpdateCityCode = `Update geography.City set code_City = $1, updated_at = now()::timestamp where id_region = $1`

var UpdateCityAll = `Update geography.City set code_City = $2,name_city = $3,id_region = $4, updated_at = now()::timestamp where id_region = $1`

var UpdateCityReference = `Update geography.City set id_region = $2, updated_at = now()::timestamp where id_region = $1`

var ShowDeleteCity = `Select	id_City, code_City,name_City from geography.City
                                where (deleted_at +  $1) > now()::timestamp`

var ShowDeleteCityId = `Select id_City, code_City,name_City from geography.City
                                where (deleted_at +  $1) > now()::timestamp and id_City = $1`

var ExistsCityName = `SELECT EXISTS(Select * from geography.City where lower(name_City) like lower($1))`
var ExistsCityId = `SELECT EXISTS(Select * from geography.City where id_city = $1)`
var ExistsCityCode = `SELECT EXISTS(Select * from geography.City where lower(code_City) like lower($1))`
